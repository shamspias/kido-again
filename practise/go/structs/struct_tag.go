package main

import "fmt"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,-"`
}

func main() {
	user := User{
		ID:       1,
		Username: "john",
		Password: "123456",
	}
	fmt.Println("user:", user)
}
