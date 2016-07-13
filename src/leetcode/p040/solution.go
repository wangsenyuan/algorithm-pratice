package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(nums, 8)
	for _, r := range res {
		fmt.Printf("%v\n", r)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	checked := make(map[string]bool)

	check := func(path []int) bool {
		key := fmt.Sprintf("%v", path)
		if checked[key] {
			return true
		}
		checked[key] = true
		return false
	}

	var combine func(start int, target int, path []int)
	combine = func(start int, target int, path []int) {
		if target == 0 {
			//fmt.Printf("candidate: %v\n", path)
			if len(path) > 0 && !check(path) {
				res = append(res, copySlice(path))
			}
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			combine(i + 1, target-candidates[i], path)
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
