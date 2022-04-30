package p2248

import "sort"

func intersection(nums [][]int) []int {
	mem := make(map[int]int)

	for _, num := range nums {
		for _, n := range num {
			mem[n]++
		}
	}

	var res []int

	for k, v := range mem {
		if v == len(nums) {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}
