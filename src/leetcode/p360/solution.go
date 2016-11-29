package main

import "fmt"

func main() {
	nums := []int{-4, -2, 2, 4}
	a, b, c := -1, 3, 5
	result := sortTransformedArray(nums, a, b, c)
	fmt.Printf("%v", result)
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	if a == 0 {
		return transform1(nums, b, c)
	} else if a > 0 {
		return transform2(nums, a, b, c)
	} else {
		return transform3(nums, a, b, c)
	}
}

func transform3(nums []int, a, b, c int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	k := 0

	for i <= j {
		x := nums[i]*(a*nums[i]+b) + c
		y := nums[j]*(a*nums[j]+b) + c
		if x <= y {
			res[k] = x
			i++
		} else {
			res[k] = y
			j--
		}
		k++
	}
	return res
}

func transform2(nums []int, a, b, c int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	k := len(nums) - 1
	for i <= j {
		x := nums[i]*(a*nums[i]+b) + c
		y := nums[j]*(a*nums[j]+b) + c
		if x >= y {
			res[k] = x
			i++
		} else {
			res[k] = y
			j--
		}
		k--
	}

	return res
}

func transform1(nums []int, b int, c int) []int {
	res := make([]int, len(nums))
	if b >= 0 {
		for i := 0; i < len(nums); i++ {
			res[i] = c + b*nums[i]
		}
	} else {
		n := len(nums)
		for i := 0; i < len(nums); i++ {
			res[n-i-1] = b*nums[i] + c
		}
	}

	return res
}
