package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Robot struct {
	Model string
}

func (d Dog) Speak() string {
	return "Woof Woff!"
}

func (r Robot) Speak() string {
	return "Beep Beep!"
}

func MakeItTalk(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	robot := Robot{
		Model: "T-800",
	}
	MakeItTalk(robot)
	dog := Dog{
		Name: "Buddy",
	}
	MakeItTalk(dog)
}
