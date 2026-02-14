package main

import (
	"fmt"
)

func makeCoffee(order string, counter chan string) {
	// Do the work
	drink := "Hot " + order

	// Send the finished drink to the counter (Channel)
	// The arrow points INTO the channel
	counter <- drink
}

func main() {
	// 1. Build the Counter (Channel) that holds strings
	pickupCounter := make(chan string)

	// 2. Tell the barista to make coffee and put it on the 'pickupCounter'
	go makeCoffee("Latte", pickupCounter)

	// 3. You wait at the counter to receive it
	// The arrow points OUT of the channel
	servedDrink := <-pickupCounter

	fmt.Println("Manager received:", servedDrink)
}
