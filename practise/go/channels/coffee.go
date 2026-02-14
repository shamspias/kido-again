package main

import "fmt"

func makeCoffee2(order string, counter chan string) {
	drink := "Hot " + order
	counter <- drink
}

func main() {
	pickupCounter := make(chan string)
	go makeCoffee2("Latte", pickupCounter)
	servedDrink := <-pickupCounter
	fmt.Println("Manager received:", servedDrink)
}
