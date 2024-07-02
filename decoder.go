package main

func decode(ids []int, vocab map[int][]byte) string {
	tokens := make([]byte, 0, len(ids))
	for _, idx := range ids {
		tokens = append(tokens, vocab[idx]...)
	}
	return string(tokens)
}
