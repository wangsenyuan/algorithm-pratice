package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		res := solve(n)
		fmt.Printf("%d", res[0])
		for j := 1; j < n; j++ {
			fmt.Printf(" %d", res[j])
		}
		fmt.Println()
	}
}

func solve(n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = 2 * (i + 1)
	}

	return res
}
