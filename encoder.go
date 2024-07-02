package main

func convertStringsToBytes(trainString string) []int {
	stringBytes := []byte(trainString)
	tokens := make([]int, len(stringBytes))

	for i, b := range stringBytes {
		tokens[i] = int(b)
	}
	return tokens
}

func TrainEncoder(trainString string, vocabSize int) (map[int][]byte, map[[2]int]int) {
	tokens := convertStringsToBytes(trainString)
	numMerges := vocabSize - 256

	merges := make(map[[2]int]int)

	for i := 0; i < numMerges; i++ {
		stats := GetStats(tokens)
		maxPair, _ := GetMaxRecurringPairs(stats)
		idxToMerge := 256 + i
		merges[maxPair] = idxToMerge
		tokens = MergeRecurringPairs(tokens, maxPair, idxToMerge)
	}
	vocab := generatePrelimVocab(merges)
	return vocab, merges
}

func generatePrelimVocab(merges map[[2]int]int) map[int][]byte {
	vocab := make(map[int][]byte)
	for i := 0; i <= 256; i++ {
		vocab[i] = []byte{byte(i)}
	}

	vocab = extendPrelimVocab(merges, vocab)
	return vocab
}

func extendPrelimVocab(merges map[[2]int]int, vocab map[int][]byte) map[int][]byte {
	for merge, idx := range merges {
		val1 := merge[0]
		val2 := merge[1]
		vocab[idx] = append(vocab[val1], vocab[val2]...)
	}
	return vocab
}

func encode(merges map[[2]int]int, text string) []int {
	textBytes := []byte(text)
	tokens := make([]int, len(textBytes))
	for i, b := range textBytes {
		tokens[i] = int(b)
	}

	for len(tokens) <= 2 {
		stats := GetStats(tokens)
		maxPair, _ := GetMaxRecurringPairs(stats)
		mergeVal, ok := merges[maxPair]
		if !ok {
			break
		}
		tokens = MergeRecurringPairs(tokens, maxPair, mergeVal)
	}
	return tokens
}
