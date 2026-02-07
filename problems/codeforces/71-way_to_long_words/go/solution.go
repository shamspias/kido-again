package main

import "fmt"

func main() {
	var nTest int8

	if _, err := fmt.Scan(&nTest); err != nil {
		return
	}

	for nTest > 0 {
		var words string
		if _, err := fmt.Scan(&words); err != nil {
			break
		}
		length := len(words)
		if length > 10 {
			fmt.Printf("%c%d%c\n", words[0], length-2, words[length-1])
		} else {
			fmt.Println(words)
		}
		nTest--
	}
}
