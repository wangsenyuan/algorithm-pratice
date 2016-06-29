package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	cnts := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		c := 0
		for j, v := range cnts {
			if a%nums[j] == 0 {
				c = max(c, v)
			}
		}
		cnts = append(cnts, c+1)
	}

	x := -1
	c := 0
	for i, v := range cnts {
		if v > c {
			c = v
			x = i
		}
	}

	result := make([]int, c, c)
	//fmt.Printf("%v\n", cnts)
	for j := x; j >= 0; j-- {
		if j < x && (nums[x]%nums[j] != 0 || cnts[x] != cnts[j]+1) {
			continue
		}
		c--
		result[c] = nums[j]
		x = j
	}

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	fmt.Printf("%v\n", largestDivisibleSubset([]int{1, 2, 4, 8, 16}))
}
