package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d/%d", &a, &b)
	c := gcd(a, b)
	a = a / c
	b = b / c
	fmt.Printf("%d/%d\n", a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
