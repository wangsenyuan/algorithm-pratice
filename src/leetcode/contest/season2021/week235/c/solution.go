package c

import "sort"

const MOD = 1000000007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	var sum int64
	pairs := make([][]int, n)

	for i := 0; i < n; i++ {
		sum += int64(abs(nums1[i] - nums2[i]))
		pairs[i] = []int{nums1[i], nums2[i]}
	}

	best := sum

	replace := func(p []int, i int) {
		tmp := sum - int64(abs(p[0]-p[1]))
		tmp += int64(abs(nums1[i] - p[1]))
		if best > tmp {
			best = tmp
		}
	}

	sort.Ints(nums1)

	for i := 0; i < n; i++ {
		cur := pairs[i]

		j := sort.Search(n, func(j int) bool {
			return nums1[j] >= cur[1]
		})

		if j < n {
			replace(cur, j)
		}
		if j > 0 {
			replace(cur, j-1)
		}
	}
	return int(best % MOD)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
