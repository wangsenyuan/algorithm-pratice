package p2806

import "sort"

func minimumTime(nums1 []int, nums2 []int, x int) int {
	// 假设t秒后，满足条件
	// 此时sum(nums2) * t + sum(nums1) - S <= x
	// 在某一时刻，将i归零的操作, 相当于减去 nums2[i] * t + nums1[i]
	n := len(nums1)

	type Pair struct {
		first  int
		second int
	}

	ps := make([]Pair, n)
	sum := make([]int, 2)
	for i := 0; i < n; i++ {
		ps[i] = Pair{nums1[i], nums2[i]}
		sum[0] += nums1[i]
		sum[1] += nums2[i]
	}

	// 假设在t秒后，完成了目标，选择了p1, p2, p3... pt
	// 那么存在nums2[p1] <= nums2[p2] <= ... nums2[pt]

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].second < ps[j].second
	})

	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			dp[j] = max(dp[j], dp[j-1]+j*ps[i].second+ps[i].first)
		}
	}

	for i := 0; i <= n; i++ {
		tmp := sum[1]*i + sum[0] - dp[i]
		if tmp <= x {
			return i
		}
	}

	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
