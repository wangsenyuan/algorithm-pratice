package main

import (
	"fmt"
	"math"
)

func smallestRange(nums [][]int) []int {
	minx, miny := 0, math.MaxInt32

	next := make([]int, len(nums))
	flag := false

	for i := 0; i < len(nums) && !flag; i++ {
		for j := 0; j < len(nums[i]) && !flag; j++ {
			min_i, max_i := 0, 0

			for k := 0; k < len(nums); k++ {
				if nums[k][next[k]] < nums[min_i][next[min_i]] {
					min_i = k
				}

				if nums[k][next[k]] > nums[max_i][next[max_i]] {
					max_i = k
				}
			}

			if miny-minx > nums[max_i][next[max_i]]-nums[min_i][next[min_i]] {
				miny = nums[max_i][next[max_i]]
				minx = nums[min_i][next[min_i]]
			}
			next[min_i]++
			flag = next[min_i] == len(nums[min_i])
		}
	}
	return []int{minx, miny}

}

func main() {
	//nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	nums := [][]int{{-5, -4, -3, -2, -1}, {1, 2, 3, 4}}
	fmt.Println(smallestRange(nums))
}
