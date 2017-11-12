package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(optimalDivision([]int{1000, 100, 10, 2}))
}

func optimalDivision(nums []int) string {
	n := len(nums)

	if n == 1 {
		return strconv.Itoa(nums[0])
	}

	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	var res = ""

	res += strconv.Itoa(nums[0]) + "/("

	for i := 1; i < n; i++ {
		res += strconv.Itoa(nums[i])
		if i < n-1 {
			res += "/"
		}
	}
	res += ")"

	return res
}
