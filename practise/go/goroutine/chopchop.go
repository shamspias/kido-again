package main

import (
	"fmt"
	"time"
)

func chopOnions() {
	for i := 0; i < 3; i++ {
		fmt.Println("🔪 Chopping onion")
		time.Sleep(100 * time.Millisecond)
	}
}

func boilPasta() {
	for i := 0; i < 3; i++ {
		fmt.Println("🍝 Boiling pasta")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 1. Summon a minion to chop onions in the background
	go chopOnions()

	// 2. YOU (the main program) boil pasta immediately
	boilPasta()

	// Note: We sleep here briefly to ensure the minion finishes before the shop closes
	time.Sleep(500 * time.Millisecond)
	fmt.Println("✅ Dinner is ready!")
}
