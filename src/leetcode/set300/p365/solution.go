package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(2, 6, 4))
	fmt.Println(canMeasureWater(2, 6, 5))
	fmt.Println(canMeasureWater(104579, 104593, 12444))
	fmt.Println(canMeasureWater(22003, 31237, 22004))
}

func canMeasureWater(x int, y int, z int) bool {
	if x == 0 && y == 0 {
		return z == 0
	}

	if z > x+y {
		return false
	}

	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
