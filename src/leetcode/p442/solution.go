package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0, len(nums))

	for _, x := range nums {
		for x <= 0 {
			x = -x
		}

		if nums[x - 1] < 0 {
			res = append(res, x)
		}

		nums[x - 1] = -nums[x - 1]
	}

	return res
}
