package main

import (
	"fmt"
)

func main() {
	diff, minerCount, blockCount, threadCount := loopInput()
	fmt.Println("difficulty:", diff)
	fmt.Println("miners:", minerCount)
	fmt.Println("rounds:", blockCount)
	fmt.Println("threads:", threadCount)
	nonce, hash := findNonce("hteh", 1)
	fmt.Printf("%x %x \n", nonce, hash)

}

/*
	function loops askInput() until correct input is submitted
	@output 4 integers corresponding to user input for difficulty level, miner count, block count, and thread count
*/
func loopInput() (int, int, int, int) {
	needInput := true
	input := []int{0, 0, 0, 0}
	for needInput {
		diff, miners, rounds, procsNum, updateNeedInput := askInput()
		input[0] = diff
		input[1] = miners
		input[2] = rounds
		input[3] = procsNum
		needInput = updateNeedInput
	}
	return input[0], input[1], input[2], input[3]
}

/*
	@output 4 integers corresponding to input accepted from user
			1 boolean representing whether the askInput() needs to be called again
*/
func askInput() (int, int, int, int, bool) {
	fmt.Println("This program will simulate a blockchain by initializing several miners and a single logger.")
	fmt.Println("The miners will attempt to solve cryptographic puzzles according to a difficulty you set.")
	fmt.Println("The difficulty is determined by comparing the most significant bits of the two hashes.")
	fmt.Println("-------------------------------------------------------------------------------")

	fmt.Println("Please choose how many leading bits you would like to be compared.")
	var difficulty int
	_, errD := fmt.Scanln(&difficulty)
	if errD != nil {
		fmt.Println("Invalid difficulty level! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Please input the number of miners you would like to simulate.")
	var numOfMiners int
	_, errM := fmt.Scanln(&numOfMiners)
	if errM != nil {
		fmt.Println("Invalid number of miners! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Please input the number of blocks you would like to add to the blockchain.")
	var numOfRounds int
	_, errR := fmt.Scanln(&numOfRounds)
	if errR != nil {
		fmt.Println("Invalid number of miners! Try again using an integer.")
		return 0, 0, 0, 0, true
	}
	fmt.Println("Please input the number of concurrent threads you would like to use.")
	var numOfProcs int
	_, errP := fmt.Scanln(&numOfProcs)
	if errP != nil {
		fmt.Println("Invalid number of threads! Try again using an integer.")
		return 0, 0, 0, 0, true
	}

	fmt.Println("Thanks! We will start the simulation with", numOfMiners, "miners on difficulty level", difficulty,
		"for", numOfRounds, "rounds.")
	fmt.Println("-------------------------------------------------------------------------------")
	return difficulty, numOfMiners, numOfRounds, numOfProcs, false
}
