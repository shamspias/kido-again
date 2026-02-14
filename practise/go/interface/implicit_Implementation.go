package main

import "fmt"

func PrintAnything(v any) {
	fmt.Println("I got:", v)
}

type Dog struct {
	Name string
}

func main() {
	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(Dog{Name: "Rex"})
}
