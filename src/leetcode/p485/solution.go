package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	var res int
	a, i := 0, 1

	for i <= len(nums) {
		if i == len(nums) || nums[i] != nums[a] {
			if nums[a] == 1 && i-a > res {
				res = i - a
			}
			a = i
		}

		i++
	}
	return res
}
