package main

import (
	"fmt"
	"sync"
	"time"
)

type BillingSystem struct {
	// Mutex 1: Protects the Money
	balanceLock sync.Mutex
	balance     int

	// Mutex 2: Protects the Stock
	inventoryLock sync.Mutex
	inventory     map[string]int
}

// 1. Update Balance (Uses Lock 1)
func (bs *BillingSystem) ChargeUser(amount int) {
	bs.balanceLock.Lock()
	defer bs.balanceLock.Unlock()

	// CRITICAL SECTION FOR MONEY
	fmt.Println("💰 Charging user...")
	time.Sleep(100 * time.Millisecond) // Simulate work
	bs.balance -= amount
}

// 2. Check Stock (Uses Lock 2)
func (bs *BillingSystem) BuyItem(item string) {
	bs.inventoryLock.Lock()
	defer bs.inventoryLock.Unlock()

	// CRITICAL SECTION FOR INVENTORY
	fmt.Println("📦 Checking stock for:", item)
	time.Sleep(100 * time.Millisecond)

	if bs.inventory[item] > 0 {
		bs.inventory[item]--
		fmt.Println("✅ Sold:", item)
	} else {
		fmt.Println("❌ Out of stock:", item)
	}
}

func main() {
	shop := BillingSystem{
		balance:   1000,
		inventory: map[string]int{"Laptop": 5},
	}

	// Because we used 2 different locks, these two actions
	// can happen AT THE SAME TIME without waiting for each other!
	go shop.ChargeUser(100)
	go shop.BuyItem("Laptop")

	time.Sleep(1 * time.Second)
}
