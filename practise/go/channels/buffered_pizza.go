package main

import "fmt"

func main() {
	// Create a channel with capacity for 2 items
	orders := make(chan string, 2)

	// I can put 2 items in without anyone listening!
	orders <- "Pizza"
	orders <- "Soda"

	// If I try to put a 3rd item in, I will BLOCK (freeze)
	// until someone takes one out.

	fmt.Println("I added two orders without a receiver waiting!")
}
