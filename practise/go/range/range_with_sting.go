package main

import "fmt"

func main() {
	word := "Here is my word"

	for index, char := range word {
		fmt.Printf("Index %d: %c\n", index, char)
	}
}
