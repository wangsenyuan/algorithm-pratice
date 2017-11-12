package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		x := num
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		}
	}
	var ans []int
	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
