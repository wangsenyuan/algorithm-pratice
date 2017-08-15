package main

import (
	"strconv"
	"fmt"
)

func fastPow(a, n int) int {
	if n == 0 {
		return 1
	}

	b := fastPow(a, n>>1)
	c := b * b
	if n&1 == 1 {
		return c * a
	}
	return c
}

func DigPow(n, p int) int {
	str := strconv.Itoa(n)
	m := len(str)

	var sum int
	for i := 0; i < m; i++ {
		x := int(str[i] - '0')
		sum += fastPow(x, p+i)
	}

	if sum%n == 0 {
		return sum / n
	}
	return -1
}


func main() {
	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(92, 1))
	fmt.Println(DigPow(695, 2))
	fmt.Println(DigPow(46288, 3))
}