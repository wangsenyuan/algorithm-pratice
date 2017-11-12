package main

import (
	"fmt"
	"sort"
)

func main() {
	rs := subsetsWithDup([]int{9, 0, 3, 5, 7})
	for _, r := range rs {
		fmt.Printf("%v\n", r)
	}
	fmt.Printf("total %d\n", len(rs))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var rs [][]int
	var subsets func(i int, set []int, lastPick int)

	subsets = func(i int, set []int, lastPick int) {
		if i >= len(nums) {
			rs = append(rs, copySlice(set))
			return
		}

		if lastPick < 0 || nums[i] != nums[lastPick] {
			subsets(i+1, set, lastPick)
		}

		subsets(i+1, append(set, nums[i]), i)
	}

	subsets(0, []int{}, -1)

	//sort.Sort(byNumber(rs))
	return rs
}

func copySlice(nums []int) []int {
	a := make([]int, len(nums))
	copy(a, nums)
	return a
}

type byNumber [][]int

func (x byNumber) Len() int {
	return len(x)
}

func (x byNumber) Less(i, j int) bool {
	a, b := x[i], x[j]
	for k := 0; k < len(a) && k < len(b); k++ {
		if a[k] < b[k] {
			return true
		}
		if a[k] > b[k] {
			return false
		}
	}

	return len(a) < len(b)
}

func (x byNumber) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
