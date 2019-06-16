package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		index := citations[mid]
		if index >= n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return n - right - 1
}
