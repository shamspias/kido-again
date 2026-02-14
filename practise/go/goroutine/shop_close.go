package main

import (
	"fmt"
	"sync" // "sync" stands for Synchronization
)

// 1. Create the checklist
var checklist sync.WaitGroup

func makeCoffee(typeOfCoffee string) {
	// 3. When this function finishes, cross 1 off the list
	defer checklist.Done()

	fmt.Println("☕ Starting:", typeOfCoffee)
}

func main() {
	// 2. We are adding 2 tasks to the checklist
	checklist.Add(2)

	go makeCoffee("Cappuccino")
	go makeCoffee("Americano")

	fmt.Println("Manager: I'm waiting for the staff to finish...")

	// 4. Block here. Do not exit the program until checklist is 0.
	checklist.Wait()

	fmt.Println("Manager: Shop is closed!")
}
