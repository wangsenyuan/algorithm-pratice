package p1926

import "sort"

const H = 17

func findMaximums(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		if p > 0 {
			left[i] = stack[p-1] + 1
		} else {
			left[i] = 0
		}
		stack[p] = i
		p++
	}
	p = 0
	for i := n - 1; i >= 0; i-- {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		if p > 0 {
			right[i] = stack[p-1] - 1
		} else {
			right[i] = n - 1
		}
		stack[p] = i
		p++
	}

	ll := make([]int, n)
	pp := make([]Pair, n)
	for i := 0; i < n; i++ {
		ll[i] = right[i] - left[i] + 1
		pp[i] = Pair{nums[i], i}
	}

	sort.Slice(pp, func(i, j int) bool {
		return pp[i].first > pp[j].first
	})

	ans := make([]int, n)

	for i, j := 0, 0; i < n; i++ {
		for j < n && ll[pp[j].second] < i+1 {
			j++
		}
		ans[i] = pp[j].first
	}
	return ans
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
