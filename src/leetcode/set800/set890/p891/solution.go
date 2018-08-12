package p891

func superEggDrop(K int, N int) int {

	f := func(x int) int {
		var ans int
		r := 1
		for i := 1; i <= K; i++ {
			r *= x - i + 1
			r /= i
			ans += r
			if ans >= N {
				return ans
			}
		}
		return ans
	}

	lo, hi := 1, N

	for lo < hi {
		mid := (lo + hi) >> 1
		if f(mid) >= N {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func superEggDrop2(K int, N int) int {
	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = i
	}

	for k := 2; k <= K; k++ {
		dp2 := make([]int, N+1)

		x := 1
		for n := 1; n <= N; n++ {
			for x < n && max(dp[x-1], dp2[n-x]) > max(dp[x], dp2[n-x-1]) {
				x++
			}
			dp2[n] = 1 + max(dp[x-1], dp2[n-x])
		}

		dp = dp2
	}

	return dp[N]
}

func superEggDrop1(K int, N int) int {

	var loop func(k, n int) int

	cache := make(map[int]int)

	loop = func(K, N int) int {
		if N == 0 {
			return 0
		}
		if K == 1 {
			return N
		}

		if v, found := cache[N*100+K]; found {
			return v
		}

		lo, hi := 1, N

		for lo+1 < hi {
			mid := (lo + hi) >> 1
			a := loop(K-1, mid-1)
			b := loop(K, N-mid)
			if a < b {
				lo = mid
			} else if a > b {
				hi = mid
			} else {
				lo, hi = mid, mid
			}
		}
		ans := 1 + min(
			max(loop(K-1, lo-1), loop(K, N-lo)),
			max(loop(K-1, hi-1), loop(K, N-hi)))
		cache[N*100+K] = ans
		return ans
	}

	return loop(K, N)
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
