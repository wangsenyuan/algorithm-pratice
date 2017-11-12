package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 0, 1, 0}
	fmt.Println(findMaxLength(nums))
}

func findMaxLength(nums []int) int {
	sum := make(map[int]int)
	pre := 0
	var res int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		pre += nums[i]
		if j, ok := sum[pre]; ok {
			if i-j > res {
				res = i - j
			}
		} else {
			sum[pre] = i
		}

		if pre == 0 {
			res = i + 1
		}
	}

	return res
}
