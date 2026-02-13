package main

import "fmt"

func main() {
	toys := map[string]int{"Teddy Bear": 1, "Latim": 2, "yo yo": 3}
	for key, value := range toys {
		fmt.Println(key, ":", value)
	}
}
