package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input a string to get numerical representation")

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()

	userInput := reader.Text()

	stringBytes := []byte(userInput)

	tokens := make([]int, len(stringBytes))
	for i, b := range stringBytes {
		tokens[i] = int(b)
	}

	// fmt.Printf("User Input: %s\n", userInput)
	// fmt.Printf("Integer representation %v\n", len(tokens))

	stats := GetStats(tokens)

	maxPair, _ := GetMaxRecurringPairs(stats)

	fmt.Printf("Most Occuring Pair: %v", maxPair)
}
