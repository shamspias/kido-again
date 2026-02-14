package main

import "fmt"

// Database 1. The Interface (The Requirement)
type Database interface {
	Save(data string) // Any worker MUST have this exact method
}

// FileSystem 2. Your Struct (The Worker)
type FileSystem struct {
	Filename string
}

// Save : You must implement the Save method!
// Once you write this, FileSystem officially implements Database.
func (fs FileSystem) Save(data string) {
	fmt.Println("Opening file:", fs.Filename)
	fmt.Println("Writing data:", data)
	fmt.Println("File closed.")
}

// SaveData  The function that demands a Database
func SaveData(db Database) {
	// Because 'db' is a Database, we KNOW it has a Save method.
	db.Save("My Important Data")
}

func main() {
	// Create the struct
	myFile := FileSystem{Filename: "backup.txt"}

	// Pass it in! It works because myFile has the Save() method.
	SaveData(myFile)
}
