package main

import "fmt"

func main() {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	for _, ans := range res {
		fmt.Println(ans)
	}
}

func findSubsequences(nums []int) [][]int {

	var res [][]int

	var find func(list []int, at int)

	find = func(list []int, at int) {
		if len(list) > 1 {
			res = append(res, doCopy(list))
		}

		used := make(map[int]bool)
		for i := at + 1; i < len(nums); i++ {
			if _, yes := used[nums[i]]; yes {
				continue
			}

			if at < 0 || nums[i] >= nums[at] {
				used[nums[i]] = true
				list = append(list, nums[i])
				find(list, i)
				list = list[:len(list) - 1]
			}
		}

	}

	find([]int{}, -1)

	return res
}

func doCopy(list []int) []int {
	dst := make([]int, len(list))
	copy(dst, list)
	return dst
}
