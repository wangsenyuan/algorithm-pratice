package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	//find the right most bit, that make m = n
	at := uint(0)
	for m != n {
		m >>= 1
		n >>= 1
		at++
	}
	return int(uint(n) << at)
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}
