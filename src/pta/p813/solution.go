package main

import "fmt"

func main() {
	var a1, b1, a2, b2 int

	fmt.Scanf("%d/%d %d/%d", &a1, &b1, &a2, &b2)
	// fmt.Printf("%d %d %d %d", a1, b1, a2, b2)
	a := a1*b2 + a2*b1
	b := b1 * b2

	c := gcd(a, b)

	a = a / c
	b = b / c

	if b == 1 {
		fmt.Printf("%d\n", a)
	} else {
		fmt.Printf("%d/%d\n", a, b)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
