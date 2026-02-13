package main

import "fmt"

func main() {
	toys := []string{"Teddy Bear", "Latim", "yo yo"}
	for _, toy := range toys {
		fmt.Println(toy)
	}
}
