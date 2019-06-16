package main

import (
	"sort"
	"fmt"
)

func triangleNumber(nums []int) int {
	sort.Ints(nums)

	var res int
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]

		for j := i + 1; j < len(nums)-1; j++ {
			b := nums[j]

			k := sort.Search(len(nums), func(k int) bool {
				return nums[k] >= a+b
			})

			if k > j+1 {
				res += k - j - 1
			}
		}
	}

	return res
}

func main() {
	nums := []int{2, 2, 3, 4}
	fmt.Println(triangleNumber(nums))
}
