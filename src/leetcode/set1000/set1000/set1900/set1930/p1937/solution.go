package p1937

import "math"

func maxPoints(points [][]int) int64 {
	m := len(points)
	n := len(points[0])

	dp := make([]int64, n)
	fp := make([]int64, n)
	for r := 0; r < m; r++ {
		cur := dp[0]
		fp[0] = dp[0] + int64(points[r][0])

		for i := 1; i < n; i++ {
			cur = max(cur, dp[i]+int64(i))
			fp[i] = cur - int64(i) + int64(points[r][i])
		}

		cur = dp[n-1] - int64(n-1)
		fp[n-1] = max(fp[n-1], dp[n-1]+int64(points[r][n-1]))

		for i := n - 2; i >= 0; i-- {
			cur = max(cur, dp[i]-int64(i))
			fp[i] = max(fp[i], cur+int64(i)+int64(points[r][i]))
		}

		copy(dp, fp)
	}
	var res int64

	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func maxPoints1(points [][]int) int64 {
	m := len(points)
	n := len(points[0])

	get := func(arr []int64, l, r int) int64 {
		l += n
		r += n
		var res int64 = math.MinInt64

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

	update := func(arr []int64, p int, v int64) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}
	arr := make([]int64, 2*n)
	dp := make([]int64, n)
	fp := make([]int64, n)
	// fp[i] = dp[j] - (i - j)
	for r := 0; r < m; r++ {
		for i := 0; i < len(arr); i++ {
			arr[i] = math.MinInt64
		}
		for i := 0; i < n; i++ {
			update(arr, i, dp[i]+int64(i))
			fp[i] = get(arr, 0, i+1) + int64(points[r][i]) - int64(i)
		}
		for i := 0; i < len(arr); i++ {
			arr[i] = math.MinInt64
		}
		for i := n - 1; i >= 0; i-- {
			update(arr, i, dp[i]-int64(i))
			fp[i] = max(fp[i], get(arr, i, n)+int64(points[r][i])+int64(i))
		}
		copy(dp, fp)
	}
	var res int64

	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
