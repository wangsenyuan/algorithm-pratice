package p3095

import (
	"slices"
	"sort"
)

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func sumOfPowers(nums []int, k int) int {
	n := len(nums)

	sort.Ints(nums)

	var dists []Dist
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dists = append(dists, Dist{i, j, nums[j] - nums[i]})
		}
	}

	slices.SortFunc(dists, func(a, b Dist) int {
		if a.val > b.val {
			return -1
		}
		if a.val < b.val {
			return 1
		}
		return 0
	})
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k)
	}
	reset := func(i int) {
		for x := 0; x < k; x++ {
			dp[i][x] = 0
		}
	}

	var res int

	// 50 * 50 * 50 * 50 * 50 too much
	for _, cur := range dists {
		d := cur.val
		// 前面在保证距离>d的情况下的dp
		// 需要知道的是，在x前面的如果选中u个数字时的计数（满足条件）
		// 后面在保证距离>=d的情况下的fp
		L := make([]int, k)
		L[0] = 1
		for l, r := cur.x, cur.x; l >= 0; l-- {
			reset(l)
			for r-1 > l && nums[r-1]-nums[l] > d {
				r--
			}
			if nums[r]-nums[l] > d {
				for x := 1; x < k; x++ {
					L[x] = add(L[x], dp[r][x-1])
				}
			}
			copy(dp[l], L)
		}

		R := make([]int, k)
		R[0] = 1
		for l, r := cur.y, cur.y; r < n; r++ {
			reset(r)
			for l+1 < r && nums[r]-nums[l+1] >= d {
				l++
			}
			if nums[r]-nums[l] >= d {
				for x := 1; x < k; x++ {
					R[x] = add(R[x], dp[l][x-1])
				}
			}
			copy(dp[r], R)
		}
		for a := 0; a <= k-2; a++ {
			b := k - 2 - a
			res = add(res, mul(d, mul(L[a], R[b])))
		}
	}

	return res
}

type Dist struct {
	x   int
	y   int
	val int
}
