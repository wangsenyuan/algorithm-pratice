package p3002

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)

	dp := make([][]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]map[int]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make(map[int]int)
		}
	}

	var dfs func(i int, mask int, changed int) int

	dfs = func(i int, mask int, changed int) int {
		if i == n {
			return 1
		}
		if v, ok := dp[i][changed][mask]; ok {
			return v
		}
		x := int(s[i] - 'a')
		cur := 1 << x
		sz := digits(mask | cur)
		if sz > k {
			// no need to add it, and start a new seg
			dp[i][changed][mask] = dfs(i+1, cur, changed) + 1
		} else {
			// sz <= k
			// two options, not change it
			ans := dfs(i+1, mask|cur, changed)
			// change it
			if changed == 0 {
				// change it
				cnt := digits(mask)
				for v := 0; v < 26; v++ {
					if v != x && (mask>>v)&1 == 0 {
						// change to some new
						if cnt == k {
							// start a new seg
							ans = max(ans, dfs(i+1, 1<<v, 1)+1)
						} else {
							ans = max(ans, dfs(i+1, mask|(1<<v), 1))
						}
					}
				}
			}
			dp[i][changed][mask] = ans
		}

		return dp[i][changed][mask]
	}

	return dfs(0, 0, 0)
}

func digits(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func maxPartitionsAfterOperations1(s string, k int) int {
	n := len(s)
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		copy(sum[i+1], sum[i])
		sum[i+1][x]++
	}

	find := func(start int, pos int, c int) int {
		l, r := start, n
		for l < r {
			mid := (l + r + 1) / 2
			var cnt int
			for j := 0; j < 26; j++ {
				t := sum[mid][j] - sum[start-1][j]
				if start <= pos && pos <= mid {
					if int(s[pos-1]-'a') == j {
						t--
					}
					if c == j {
						t++
					}
				}
				if t > 0 {
					cnt++
				}
			}
			if cnt <= k {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return r
	}

	f := make([]int, n+2)
	for i := n; i > 0; i-- {
		f[i] = f[find(i, 0, 0)+1] + 1
	}

	var vec []int
	for now := 0; now < n; {
		vec = append(vec, now+1)
		now = find(now+1, 0, 0)
	}

	ans := len(vec) - 1

	for i := 1; i <= n; i++ {
		idx := search(len(vec), func(j int) bool {
			return vec[j] > i
		})
		idx--
		for j := 0; j < 26; j++ {
			now := vec[idx]
			var cnt int
			for now <= i {
				cnt++
				now = find(now, i, j) + 1
			}
			ans = max(ans, idx+cnt+f[now])
		}
	}

	return ans

}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
