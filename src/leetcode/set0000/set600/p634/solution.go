package main

import "fmt"

func findDerangement(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 0
	}

	// n - 2, n - 1
	a, b := 1, 0

	mod := 1000000007

	for i := 2; i <= n; i++ {
		c := ((i - 1) * (b + a)) % mod
		a, b = b, c
	}

	return b
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d => %d\n", i, findDerangement(i))
	}
}
