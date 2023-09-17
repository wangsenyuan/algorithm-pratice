package p2863

func maximumSum(nums []int) int64 {
	n := len(nums)
	lf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if lf[i] == 0 {
			for j := i; j <= n; j += i {
				if lf[j] == 0 {
					lf[j] = i
				}
			}
		}
	}

	dp := make(map[int]int64)

	for i, num := range nums {
		tmp := i + 1
		cnt := make(map[int]int)
		for tmp > 1 {
			cnt[lf[tmp]]++
			tmp /= lf[tmp]
		}
		tmp = 1
		for k, v := range cnt {
			if v%2 == 1 {
				tmp *= k
			}
		}
		dp[tmp] += int64(num)
	}
	var best int64
	for _, v := range dp {
		if v > best {
			best = v
		}
	}
	return best
}
