package main

import "fmt"

func main() {
	s := 11
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(s, nums))
}

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	res := len(nums) + 1

	j := 0
	for i := 0; i < len(nums); i++ {
		for j < len(nums) && sum < s {
			sum += nums[j]
			j++
		}
		if sum < s {
			break
		}
		if j-i < res {
			res = j - i
		}
		sum -= nums[i]
	}

	if res > len(nums) {
		return 0
	}
	return res
}
