package main

import "fmt"

func CreateCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	counterA := CreateCounter()
	counterB := CreateCounter()

	fmt.Println("Counter A: ", counterA())
	fmt.Println("Counter A: ", counterA())
	fmt.Println("Counter B: ", counterB())
}
