package main

import "fmt"

func main() {
	res := combinationSum3(3, 9)

	for _, r := range res {
		fmt.Println(r)
	}
}

func combinationSum3(k int, n int) [][]int {

	var sum func(x int, n int, nums []int, j int)

	var res [][]int

	sum = func(x int, m int, nums []int, j int) {
		if m == 0 && j == k {
			res = append(res, duplicate(nums, j))
			return
		}

		if m == 0 || j > k {
			return
		}

		for i := 1; i+x <= m && i+x < 10; i++ {
			nums[j] = i + x
			sum(x+i, m-x-i, nums, j+1)
		}
	}

	nums := make([]int, 10)
	sum(0, n, nums, 0)
	return res
}

func duplicate(nums []int, l int) []int {
	cp := make([]int, l)
	copy(cp, nums)
	return cp
}
