package main

import (
	"fmt"
	"sync"
	"time"
)

// The Barista Function
func makeCoffee(id int, counter chan string, wg *sync.WaitGroup) {
	// 3. When done, tell the Supervisor (WaitGroup) "I'm finished!"
	defer wg.Done()

	fmt.Printf("👷 Barista %d: Making coffee...\n", id)
	time.Sleep(time.Second) // simulating work

	drink := fmt.Sprintf("Coffee from Barista %d", id)

	// 4. Put drink on the counter
	counter <- drink
}

func main() {
	// 1. Setup the Shop
	counter := make(chan string) // The channel
	var wg sync.WaitGroup        // The checklist

	// 2. Hire 3 Baristas
	numBaristas := 3
	for i := 1; i <= numBaristas; i++ {
		wg.Add(1)                      // Add to checklist
		go makeCoffee(i, counter, &wg) // Pass the checklist pointer
	}

	// THE TRICK: The "Shift Supervisor"
	// We start a background routine just to watch the checklist.
	// Once everyone is done, it closes the channel.
	go func() {
		wg.Wait() // Wait for all 3 baristas to finish
		fmt.Println("🔒 Supervisor: All baristas done. Channel closed.")

		close(counter) // Close the channel (Shop is closed!)
	}()

	// 5. The Manager (Main) picks up drinks
	// This loop keeps running until the channel is CLOSED by the Supervisor.
	fmt.Println("Manager: Waiting for coffees...")

	for drink := range counter {
		fmt.Println("✅ Received:", drink)
	}

	fmt.Println("Manager: Work is done, going home!")
}
