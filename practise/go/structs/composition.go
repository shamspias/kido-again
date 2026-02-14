package main

import "fmt"

type Player struct {
	Username string
	Score    int
	IsActive bool
}

type Admin struct {
	Player
	AdminLevel int
}

func main() {
	boss := Admin{
		Player:     Player{Username: "SuperBoss", Score: 99, IsActive: true},
		AdminLevel: 5,
	}

	fmt.Println(boss.Username)
}
