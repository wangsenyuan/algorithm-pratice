package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		var c = a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

func main() {
	//fmt.Printf("%d + %d = %d\n", 1, 2, getSum(1, 2))
	fmt.Printf("%d + %d = %d\n", 2, 3, getSum(3, 2))
}
