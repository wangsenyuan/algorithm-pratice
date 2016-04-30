package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}

	for i := uint(0); i < 64; i += 2 {
		x := 1 << i
		if num&x > 0 {
			y := ^x
			if num&y == 0 {
				return true
			}
		}
	}

	return false
}

func main() {
	for i := 1; i < math.MaxInt64 && i > 0; i *= 4 {
		fmt.Printf("%d is %v\n", i, isPowerOfFour(i))
	}
}
