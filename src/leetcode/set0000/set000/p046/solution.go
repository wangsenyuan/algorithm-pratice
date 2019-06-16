package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	rs := permute(nums)
	for _, r := range rs {
		fmt.Printf("%v\n", r)
	}
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	x := nums[0]

	ps := permute(nums[1:])
	rs := make([][]int, 0, len(ps))
	for _, p := range ps {
		for i := 0; i < len(nums); i++ {
			t := make([]int, i, len(nums))
			copy(t, p[0:i])
			t = append(t, x)
			t = append(t, p[i:]...)
			rs = append(rs, t)
		}
	}

	return rs
}
