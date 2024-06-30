package main

import "math"

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
