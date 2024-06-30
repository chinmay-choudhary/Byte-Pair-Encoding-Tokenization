package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Input a string to get numerical representation")

	// reader := bufio.NewScanner(os.Stdin)
	// reader.Scan()

	// userInput := reader.Text()

	stringBytes := []byte(LONGTEXT)

	tokens := make([]int, len(stringBytes))
	for i, b := range stringBytes {
		tokens[i] = int(b)
	}

	// fmt.Printf("User Input: %s\n", userInput)
	// fmt.Printf("Integer representation %v\n", tokens)
	// fmt.Printf("Len of tokens before merging recurring tokens: %v\n", len(tokens))
	vocabSize := 276
	numMerges := vocabSize - 256
	ids := tokens

	for i := 0; i < numMerges; i++ {
		stats := GetStats(ids)
		maxPair, _ := GetMaxRecurringPairs(stats)
		idxToMerge := 256 + i
		ids = MergeRecurringPairs(ids, maxPair, idxToMerge)
	}
	fmt.Println("Original token length: ", len(tokens))
	fmt.Println("Merged ids length: ", len(ids))
	compressionRatio := float64(len(tokens)) / float64(len(ids))
	fmt.Printf("Compression Ratio: %.2f", compressionRatio)
}
