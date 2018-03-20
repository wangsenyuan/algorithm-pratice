package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(solve(n, m))
		t--
	}
}

func solve(n int, m int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return m
	}
	return n - 1 + 2*(m-1)
}
