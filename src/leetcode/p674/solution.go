package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	var ans int

	for i, j := 1, 0; i <= len(nums); i++ {
		if i < len(nums) && nums[i] > nums[i-1] {
			continue
		}

		if i-j > ans {
			ans = i - j
		}
		j = i
	}

	return ans
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))

}
