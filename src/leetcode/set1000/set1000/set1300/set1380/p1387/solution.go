package p1387

import "sort"

var ps []int

func init() {
	ps = make([]int, 1001)
	ps[1] = 0

	var loop func(x int) int

	loop = func(x int) int {
		if x == 1 {
			return 0
		}

		if x&1 == 1 {
			return loop(3*x+1) + 1
		} else {
			return loop(x/2) + 1
		}
	}

	for i := 1; i <= 1000; i++ {
		ps[i] = loop(i)
	}
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first || this[i].first == this[j].first && this[i].second < this[j].second
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func getKth(lo int, hi int, k int) int {
	var nums []Pair

	for i := lo; i <= hi; i++ {
		nums = append(nums, Pair{ps[i], i})
	}

	sort.Sort(Pairs(nums))

	return nums[k-1].second
}
