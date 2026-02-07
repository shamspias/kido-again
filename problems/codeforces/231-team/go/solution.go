package main

import "fmt"

func main() {
	var nTest int16
	var solveCount uint16 = 0

	if _, err := fmt.Scan(&nTest); err != nil {
		return
	}

	for nTest > 0 {
		var numberOne, numberTwo, numberThree int8

		if _, err := fmt.Scan(&numberOne, &numberTwo, &numberThree); err != nil {
			break
		}
		if numberOne+numberTwo+numberThree > 1 {
			solveCount++
		}
		nTest--
	}
	fmt.Println(solveCount)
}
