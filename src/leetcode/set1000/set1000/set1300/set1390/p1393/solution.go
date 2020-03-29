package p1393

import "sort"

type Pair struct {
	first  int64
	second int64
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	a := this[i].first - this[i].second
	b := this[j].first - this[j].second
	return a < b
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
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

func findLucky(arr []int) int {
	sort.Ints(arr)
	n := len(arr)

	var ans int = -1

	var j int

	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			cnt := i - j
			if cnt == arr[i-1] {
				ans = arr[i-1]
			}
			j = i
		}
	}

	return ans
}
