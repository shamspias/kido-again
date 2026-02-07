package main

import "fmt"

func main() {
	var melonWeight int8
	n, err := fmt.Scan(&melonWeight)
	if err != nil {
		return
	}
	if n == 1 && melonWeight > 2 && melonWeight&1 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
