package main

import (
	"fmt"
	"time"
)

func makeCoffee(typeOfCoffee string) {
	fmt.Println("☕ Barista started making:", typeOfCoffee)
	// Pretend this takes 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("✅ Barista finished:", typeOfCoffee)
}

func main() {
	// You (the Main Manager) tell a barista: "Go make a latte!"
	go makeCoffee("Latte")

	// You immediately take the next order, you don't wait for the Latte!
	fmt.Println("You: Can I help the next customer?")

	// IMPORTANT: If you (Main) go home, the shop closes.
	// We wait 2 seconds here just so the Barista has time to finish.
	time.Sleep(2 * time.Second)
}
