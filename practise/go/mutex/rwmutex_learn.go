package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	menu     = "Pizza"
	menuLock sync.RWMutex // Note the 'RW'
)

// The Customers (Readers)
func readMenu(id int) {
	// 1. Get a "Read Lock"
	// Many customers can hold this AT THE SAME TIME
	menuLock.RLock()
	defer menuLock.RUnlock()

	fmt.Printf("👀 Customer %d sees: %s\n", id, menu)
	time.Sleep(1 * time.Second) // Taking time to read
}

// The Manager (Writer)
func changeMenu(newItem string) {
	// 2. Get a "Write Lock"
	// This BLOCKS until all readers are gone.
	// No one can read while this is held.
	menuLock.Lock()
	defer menuLock.Unlock()

	fmt.Println("\n👨‍🍳 MANAGER: Changing menu to", newItem)
	menu = newItem
	time.Sleep(1 * time.Second) // Taking time to write
	fmt.Println("👨‍🍳 MANAGER: Done changing menu.\n")
}

func main() {
	// 1. Let 5 customers read simultaneously
	// They will all print almost instantly because RLock allows sharing.
	for i := 1; i <= 5; i++ {
		go readMenu(i)
	}

	time.Sleep(200 * time.Millisecond)

	// 2. Manager tries to change it
	// He has to wait for the first 5 readers to finish!
	go changeMenu("Pasta")

	time.Sleep(2 * time.Second)

	// 3. New customers see the new menu
	go readMenu(6)

	time.Sleep(1 * time.Second)
}
