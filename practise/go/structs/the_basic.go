package main

import "fmt"

type Player struct {
	Username string
	Score    int
	IsActive bool
}

func main() {
	players := []*Player{
		{Username: "A", Score: 10, IsActive: true},
		{Username: "B", Score: 20, IsActive: true},
	}
	for _, player := range players {
		fmt.Println(player.Username, "has", player.Score, "points.")
	}
}
