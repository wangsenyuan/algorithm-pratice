// chapter4_execise
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	nums := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8}
	fmt.Println(UniqueInts(nums))
}

func UniqueInts(nums []int) []int {
	result := make([]int, 0, len(nums))
	checked := map[int]bool{}

	for _, num := range nums {
		if checked[num] {
			continue
		}
		checked[num] = true
		result = append(result, num)
	}

	return result
}
