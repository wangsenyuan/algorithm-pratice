package main

import (
	"fmt"
	"math"
)

func main() {
	var t int

	fmt.Scanf("%d", &t)

	for t > 0 {
		var Y int64

		fmt.Scanf("%d", &Y)

		fmt.Println(solve(Y))

		t--
	}
}

func solve(Y int64) int64 {
	var ans int64
	for B := int64(1); B <= 700 && B < Y; B++ {
		X := float64(Y - B)
		A := int64(math.Sqrt(X))
		ans += A
	}

	return ans
}
