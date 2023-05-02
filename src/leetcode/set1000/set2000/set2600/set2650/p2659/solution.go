package p2659

import "sort"

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)

	arr := make([]Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = Pair{nums[i], i}
	}

	cnt := make([]int, 2*n)
	for i := n; i < 2*n; i++ {
		cnt[i] = 1
	}
	for i := n - 1; i > 0; i-- {
		cnt[i] = cnt[2*i] + cnt[2*i+1]
	}

	set := func(p int, v int) {
		p += n
		for p > 0 {
			cnt[p] += v
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += n
		r += n
		var res int
		for l < r {
			if l&1 == 1 {
				res += cnt[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += cnt[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first < arr[j].first
	})

	var res int64

	if arr[0].second != 0 {
		res += int64(arr[0].second)
	}

	for i, cur := range arr {
		// remove operation
		res++
		// move to next position
		x := cur.second
		set(x, -1)
		if i+1 < len(arr) {
			y := arr[i+1].second
			if x < y {
				res += int64(get(x, y))
			} else {
				res += int64(get(x, n))
				res += int64(get(0, y))
			}
		}
	}

	return res
}

type Pair struct {
	first  int
	second int
}
