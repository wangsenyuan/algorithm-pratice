package p1425

import "math"

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)

	arr := make([]int, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = math.MinInt32
	}

	update := func(pos int, v int) {
		pos += n

		arr[pos] = v

		for pos > 1 {
			arr[pos>>1] = max(arr[pos^1], arr[pos])
			pos >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n
		res := math.MinInt32

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	res := math.MinInt32

	for i := 0; i < n; i++ {
		j := max(0, i-k)
		x := get(j, i)
		y := max(0, x) + nums[i]

		res = max(res, y)

		update(i, y)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
