package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Training Tokenizer")
	vocab, merges := TrainEncoder(LONGTEXT, 296)

	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Input string to encode: ")

	reader.Scan()

	userInput := reader.Text()

	ids := encode(merges, userInput)

	fmt.Printf("Tokenized Input: %v\n", ids)

	decodedUserInput := decode(ids, vocab)

	fmt.Printf("Decoded Input: %v\n", decodedUserInput)

	fmt.Printf("%v\n", userInput == decodedUserInput)
}
