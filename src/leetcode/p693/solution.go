package main

import "fmt"

func hasAlternatingBits(n int) bool {
	last := n & 1
	n = n >> 1
	for n > 0 {
		tmp := n & 1
		n = n >> 1
		if tmp == last {
			return false
		}
		last = tmp
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
}
