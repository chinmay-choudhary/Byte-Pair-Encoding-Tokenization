package main

import (
	"fmt"
	"math"
)

func GetStats(tokens []int) map[[2]int]int {
	count := make(map[[2]int]int)
	for i := 0; i < len(tokens)-1; i++ {
		currPair := [2]int{tokens[i], tokens[i+1]}
		_, ok := count[currPair]
		if !ok {
			count[currPair] = 0
		} else {
			count[currPair] += 1
		}
	}
	return count
}

func GetMaxRecurringPairs(m map[[2]int]int) ([2]int, int) {
	maxKey := [2]int{0, 0}
	maxValue := math.MinInt64 // Initialize with the smallest possible integer

	for k, v := range m {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}

	return maxKey, maxValue
}

func MergeRecurringPairs(tokens []int, recurringPair [2]int, idxToMerge int) []int {
	fmt.Printf("Merging %v into %v\n", recurringPair, idxToMerge)
	newTokens := []int{}
	i := 0
	for i < len(tokens) {
		if i < len(tokens)-1 && tokens[i] == recurringPair[0] && tokens[i+1] == recurringPair[1] {
			newTokens = append(newTokens, idxToMerge)
			i += 2
		} else {
			newTokens = append(newTokens, tokens[i])
			i++
		}
	}
	return newTokens
}
