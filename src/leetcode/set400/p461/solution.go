package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(0, 4))
}

func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for z > 0 {
		if z&1 == 1 {
			cnt++
		}
		z = z >> 1
	}

	return cnt
}
