package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 6, 7}
	res := combinationSum(nums, 7)
	for _, r := range res {
		fmt.Printf("%v\n", r)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var combine func(start int, target int, path []int)
	combine = func(start int, target int, path []int) {
		if target == 0 {
			if len(path) > 0 {
				res = append(res, copySlice(path))
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			combine(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}

	combine(0, target, make([]int, 0, 10))

	return res
}

func copySlice(a []int) []int {
	b := make([]int, len(a), len(a))
	copy(b, a)
	return b
}
