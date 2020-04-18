package maxarea

import "sort"

func maxArea(height []int) int {
	n := len(height)
	var left int
	right := n - 1

	var res int

	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func maxArea2(height []int) int {
	n := len(height)
	// height[i] * i - height[i] * j when height[i] < height[j]
	// or height[j] * i - height[j] * j when height[i] > height[j]
	// if we fix i, then we need to find left-most j, has height[j] >= height[i]
	ps := make([]Pair, n)

	for i := 0; i < n; i++ {
		ps[i] = Pair{height[i], i}
	}

	sort.Sort(Pairs(ps))

	var left = n - 1
	var right = 0

	var res int

	for i := 0; i < n; i++ {
		cur := ps[i]
		if left < cur.second {
			res = max(res, cur.first*(cur.second-left))
		}
		if right > cur.second {
			res = max(res, cur.first*(right-cur.second))
		}
		left = min(left, cur.second)
		right = max(right, cur.second)
	}

	return res
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first > ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func maxArea1(height []int) int {
	n := len(height)

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			res = max(res, min(height[i], height[j])*(i-j))
		}
	}

	return res
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
