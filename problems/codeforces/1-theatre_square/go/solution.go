package main

import "fmt"

func calculatePlace(m, n, a int64) int64 {
	mPart := (m + a - 1) / a
	nPart := (n + a - 1) / a
	return mPart * nPart
}
func main() {
	var m, n, a int64
	if _, err := fmt.Scan(&m, &n, &a); err != nil {
		return
	}
	fmt.Println(calculatePlace(m, n, a))
}
