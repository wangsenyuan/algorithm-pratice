package p2680

const H = 30

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	// 还是要贪心
	cnt := make([]int, H)
	var best int64

	for i := 0; i < n; i++ {
		// 如果高的k位一致的情况下，选择最小的那个数移动
		num := nums[i]
		for j := 0; j < H; j++ {
			cnt[j] += (num >> j) & 1
		}
		best |= int64(num)
	}

	for i := 0; i < n; i++ {
		num := nums[i]
		for j := 0; j < H; j++ {
			cnt[j] -= (num >> j) & 1
		}

		tmp := int64(num) << k

		for j := 0; j < H; j++ {
			if cnt[j] > 0 {
				tmp |= 1 << j
			}
			cnt[j] += (num >> j) & 1
		}
		best = max(best, tmp)
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
