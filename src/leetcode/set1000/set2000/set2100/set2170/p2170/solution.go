package p2170

import "sort"

func minimumOperations(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	even := make([]int, 0, (n+1)/2)
	odd := make([]int, 0, n-cap(even))

	for i := 0; i < n; i++ {
		if i&1 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	a := most(even)
	b := most(odd)
	if a[0].second != b[0].second {
		// keep most of even & odd
		return n - a[0].first - b[0].first
	}
	if len(a) == 1 && len(b) == 1 {
		// change one of them to another number
		return min(a[0].first, b[0].first)
	}
	if len(b) == 1 {
		return len(even) - a[1].first
	}
	if len(a) == 1 {
		return len(odd) - b[1].first
	}
	x := n - a[0].first - b[1].first
	y := n - a[1].first - b[0].first
	return min(x, y)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func most(arr []int) []Pair {
	sort.Ints(arr)
	res := make([]Pair, 0, 2)

	for i := 0; i < len(arr); {
		j := i
		for i < len(arr) && arr[i] == arr[j] {
			i++
		}
		cnt := i - j
		p := Pair{cnt, arr[j]}
		if len(res) == 0 || cnt > res[0].first {
			if len(res) == 2 {
				res[1] = res[0]
			}
			if len(res) == 0 {
				res = append(res, p)
			} else {
				res[0] = p
			}
		} else if len(res) == 1 || cnt > res[1].first {
			if len(res) == 1 {
				res = append(res, p)
			} else {
				res[1] = p
			}
		}
	}

	return res
}

type Pair struct {
	first, second int
}
