package main

import "fmt"

func fakeUpdate(toy string) {
	toy = "Teddy Bear" + toy
}

func realUpdate(toy *string) {
	*toy = *toy + " Teddy Bear"
}

func main() {
	toy := "Supper"
	fakeUpdate(toy)
	fmt.Println("Toy: ", toy)

	realUpdate(&toy)
	fmt.Println("Toy: ", toy)
}
