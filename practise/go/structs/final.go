package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// --- PART 1: JSON Structs ---

type User struct {
	ID       int    `json:"id"`                 // Will appear as "id"
	Username string `json:"username,omitempty"` // If empty, will vanish!
	Password string `json:"-"`                  // The "-" tells Go: HIDE THIS!
}

// --- PART 2: Memory Alignment Structs ---

// BadStruct : The fields are mixed up (Bool, Int, Bool)
// Imagine packing a suitcase badly: Small, Big, Small. Lots of gaps!
type BadStruct struct {
	IsAdmin  bool  // 1 byte + (7 bytes padding)
	Score    int64 // 8 bytes
	IsActive bool  // 1 byte + (7 bytes padding)
}

// GoodStruct : The fields are sorted (Int, Bool, Bool)
// Packed tightly: Big, Small, Small.
type GoodStruct struct {
	Score    int64 // 8 bytes
	IsAdmin  bool  // 1 byte
	IsActive bool  // 1 byte
	// (Only 6 bytes padding at the very end to finish the block)
}

func main() {
	// ==========================================
	// TEST 1: JSON Tags (Serialization)
	// ==========================================
	fmt.Println("--- JSON TEST ---")

	// Case A: A full user
	user1 := User{ID: 1, Username: "GoExpert", Password: "SecretPassword123"}

	// We use MarshalIndent to make it look pretty (like Python's json.dumps indent=2)
	jsonData1, _ := json.MarshalIndent(user1, "", "  ")

	fmt.Printf("User 1 (Full):\n%s\n", string(jsonData1))
	// Notice: Password is GONE because of `json:"-"`

	fmt.Println()

	// Case B: A user with NO username (testing omitempty)
	user2 := User{ID: 2, Password: "Hidden"}
	jsonData2, _ := json.MarshalIndent(user2, "", "  ")

	fmt.Printf("User 2 (No Username):\n%s\n", string(jsonData2))
	// Notice: "username" key is completely missing!

	// ==========================================
	// TEST 2: Memory Alignment (The "unsafe" truth)
	// ==========================================
	fmt.Println("\n--- MEMORY ALIGNMENT TEST ---")

	bad := BadStruct{}
	good := GoodStruct{}

	// unsafe.Sizeof tells us exactly how many bytes of RAM this takes
	fmt.Printf("BadStruct size:  %d bytes\n", unsafe.Sizeof(bad))
	fmt.Printf("GoodStruct size: %d bytes\n", unsafe.Sizeof(good))

	fmt.Println("\nOptimization:",
		unsafe.Sizeof(bad)-unsafe.Sizeof(good),
		"bytes saved per struct!")
}
