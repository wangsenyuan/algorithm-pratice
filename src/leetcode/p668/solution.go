package main

import (
	"fmt"
	"sort"
)

func findKthNumber(m int, n int, k int) int {
	count := func(x int) int {
		sum := 0
		for i := 1; i <= m; i++ {
			sum += min(x/i, n)
		}
		return sum
	}

	return sort.Search(m*n+1, func(x int) bool {
		return count(x) >= k
	})

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(findKthNumber(3, 3, 5))
	fmt.Println(findKthNumber(2, 3, 6))

}
