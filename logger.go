// logger funtions and struct

package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

func loggerVerify(b *Block) {

	prevBlockHash := b.hash
	nonce := b.nonce
	diff := b.difficulty

	hashSeed := bytes.NewBuffer(prevBlockHash[:]).String()
	diffSlice := make([]byte, diff)
	strNonce := strconv.Itoa(nonce)
	verifyHash := sha256.Sum256([]byte(strNonce + hashSeed))
	x := verifyHash[:diff]
	// verify nonce
	if bytes.Equal(x, diffSlice) {
		return true
	}

	return false
	// append block to list of blocks
	//end routine and start logger notify
}

func runLogger(minerArray, pushChanData newmsg, oldblock *Block) {
// close current channels with miner
// select case
block = newmsg.miner.Block 
newBlock := createBlock(block.nonce, block.hash, block.diff, &newmsg.miner.Block)
if (loggerVerify(block) == true){
	for i := 0; i < minerArray.length(); i++{
		minerArray[i].notifyChan <- true 
		minerArray[i].pullChan <- newBlock
	}
}
else{
	newmsg.miner.pullChan <- oldblock
}
 break 
}



