package main

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
)

//
type Miner struct {
	id           int
	toMinerChan  chan Block        //Channel used to receive block from logger, unique to each Miner
	toLoggerChan chan toLoggerData //Channel used to send block and miner to logger, all Miners use the same channel
	notifyChan   chan bool         //Channel used to receive a notification that there is a new block/a block has been verified.
	//Primarily used inside the findNonce function
}

type toLoggerData struct {
	miner *Miner
	block Block
}

func createMiner(id int, toLoggerChan chan toLoggerData, blockCount int) Miner {
	pull := make(chan Block, blockCount)
	notifyChan := make(chan bool, blockCount)
	miner := Miner{id, pull, toLoggerChan, notifyChan}
	return miner
}

//Main routine for Miner.
//@input miner being ran, difficulty of the hash
//If there is a new block from the logger attempt to find a correct nonce
func run(currMiner *Miner, diff int) {
	var prevBlock Block
	var nonce int
	var hash [32]byte
	for true {

		select {
		case newBlock := <-currMiner.toMinerChan:

			prevBlock = newBlock
			prevBlockHash := prevBlock.hash
			randNum := rand.Intn(100)
			if randNum < 1 {

				nonce, hash = findBadNonce(prevBlockHash, diff)
			} else {

				nonce, hash = findNonce(prevBlockHash, diff, currMiner)

			}

			//1% probability to send a bad block.

			newBlockHeight := prevBlock.height + 1 //Block's number/ID
			block := createBlock(nonce, hash, diff, &prevBlock, newBlockHeight)
			setBlockHash(&block)
			data := toLoggerData{currMiner, block}
			currMiner.toLoggerChan <- data
		default:
			//Do Nothing

		}
	}
}

//Function to find the correct nonce.
//@input: hash of previous block, difficulty of the puzzle, and current miner
//@output: nonce that produces a hash with x number of leading zeroes, that hash
func findNonce(seed [32]byte, diff int, miner *Miner) (int, [32]byte) {
	hashSeed := bytes.NewBuffer(seed[:]).String()
	diffSlice := make([]byte, diff) //Slice which is used to compare "difficulty" leading zeroes of the hashes.
	nonceFound := false
	nonce := rand.Intn(100)

	//If there is no new block:
	//This for loop continuely creates new hashes by concatenating a new nonce and the previous blocks hash each iteration.
	//It then compares it to the difficulty
	var newHash [32]byte
	for !nonceFound {
		select {
		case here := <-miner.notifyChan:
			if here == true {
				break
			}
		default:

			nonce++
			strNonce := strconv.Itoa(nonce)
			newHash = sha256.Sum256([]byte(strNonce + hashSeed))

			x := newHash[:diff] // Taking the x leading zeroes of the foundhash
			if bytes.Equal(x, diffSlice) {
				nonceFound = true
			}
		}
	}
	return nonce, newHash
}

//Function to find a bad nonce. Almost the same as findNonce, except it wants to find a nonce that produces a hash without that number of leading zeroes.
func findBadNonce(seed [32]byte, diff int) (int, [32]byte) {
	hashSeed := bytes.NewBuffer(seed[:]).String()
	diffSlice := make([]byte, diff) //Slice which is used to compare the found hash
	nonceFound := false
	nonce := -1
	var newHash [32]byte
	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash = sha256.Sum256([]byte(strNonce + hashSeed))
		x := newHash[:diff] // Taking the x leading zeroes of the foundhash
		if !bytes.Equal(x, diffSlice) {
			nonceFound = true
		}
	}
	return nonce, newHash
}
