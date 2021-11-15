package main

import (
	"bytes"
	"fmt"
	//	"sync"
	//	"time"
	"crypto/sha256"
	"strconv"
	//"strings"
)

type Block struct {
	previousBlock *Block
	nonce         int
	tx            string //Do we need a tx at all? I guess it would then need to be its own struct, so maybe unessecary

}

func main() {
	nonce := findNonce("abcd")
	fmt.Printf("%d \n", nonce)
}

func askInput() {

}

func findNonce(hash string) int {

	diffSlice := []byte{0}
	fmt.Printf("The difficulty is: %x \n", diffSlice)
	nonceFound := false
	nonce := -1
	var newHash []byte

	for !nonceFound {
		nonce++
		strNonce := strconv.Itoa(nonce)
		newHash := sha256.Sum256([]byte(strNonce + hash))
		if nonce%10000 == 0 {
			fmt.Printf("Here %x \n", newHash)
		}

		x := newHash[:1]
		if bytes.Equal(x, diffSlice) {
			fmt.Printf("In if statement %x \n", newHash)
			nonceFound = true
		}
	}
	fmt.Printf("%x \n", newHash) //So there is a problem with how outside of the for loop, newHash won't print out.
	return nonce

}
