package main

import "fmt"

type Robot struct {
	Name   string
	Energy int
}

func (r *Robot) DrainEnergy() {
	r.Energy = r.Energy - 10
	fmt.Println("Inside Method Energy is: ", r.Energy)
}

func main() {

	myBot := Robot{"Pinka", 100}
	myBot.DrainEnergy()
	fmt.Println("Back in main, energy is: ", myBot.Energy)
}
