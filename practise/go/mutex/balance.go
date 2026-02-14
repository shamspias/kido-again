package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. The Shared Resource (The Bank Account)
var balance int = 0

// 2. The Key (Mutex)
// We usually keep the data and the lock together
var lock sync.Mutex

func deposit() {
	// 3. LOCK THE DOOR!
	// If someone else is already inside, this line PAUSES and waits.
	lock.Lock()

	// --- CRITICAL SECTION (Safe Zone) ---
	// Only ONE goroutine can be here at a time.
	temp := balance
	time.Sleep(1 * time.Millisecond) // Simulating slow processing
	balance = temp + 10
	fmt.Println("💰 Deposited $10. New Balance:", balance)
	// ------------------------------------

	// 4. UNLOCK THE DOOR!
	// If you forget this, the program freezes forever (Deadlock).
	lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// Let's have 5 people deposit money at the same time
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deposit()
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance:", balance) // Should be 50
}
