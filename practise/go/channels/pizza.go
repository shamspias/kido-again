package main

import "fmt"

func cook(ordersChannel chan string) {
	// The cook waits here until an order slides down the tube
	msg := <-ordersChannel
	fmt.Println("👨‍🍳 Cook received order:", msg)
}

func main() {
	// 1. Create a channel (The Tube) that carries strings
	// 'make' is how we initialize things in Go
	kitchenTube := make(chan string)

	// 2. Start the cook minion
	go cook(kitchenTube)

	// 3. Send an order down the tube
	fmt.Println("💁 Waiter sending order...")
	kitchenTube <- "Pepperoni Pizza" // The arrow points INTO the variable

	// Wait for input so the program doesn't exit immediately
	var input string
	fmt.Scanln(&input)
}
