package p3006

const D = 60

func findMaximumNumber(k int64, x int) int64 {
	var l, r int64 = 0, (k + 1) << (x - 1)

	for l < r {
		mid := (l + r) >> 1
		if count(mid, x) > k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func count(num int64, x int) int64 {
	//计算不超过num的价值和
	num++
	var ds []int
	for tmp := num; tmp > 0; tmp >>= 1 {
		ds = append(ds, int(tmp&1))
	}
	n := len(ds)
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, D)
		for j := 0; j < D; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(i int, cnt int, eq bool) int64

	dfs = func(i int, cnt int, eq bool) int64 {
		if i < 0 {
			return int64(cnt)
		}
		var res int64
		if !eq {
			it := &dp[i][cnt]
			if *it >= 0 {
				return *it
			}
			defer func() {
				*it = res
			}()
		}

		up := 1
		if eq {
			up = ds[i]
		}

		for d := 0; d <= up; d++ {
			nc := cnt
			if d == 1 && (i+1)%x == 0 {
				nc++
			}
			res += dfs(i-1, nc, eq && d == ds[i])
		}

		return res
	}

	return dfs(n-1, 0, true)
}
