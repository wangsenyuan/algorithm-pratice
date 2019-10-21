package p1234

func balancedString(s string) int {
	n := len(s)

	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 4)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			dp[i+1][j] = dp[i][j]
		}

		if s[i] == 'Q' {
			dp[i+1][0]++
		} else if s[i] == 'W' {
			dp[i+1][1]++
		} else if s[i] == 'E' {
			dp[i+1][2]++
		} else {
			dp[i+1][3]++
		}
	}

	m := n / 4
	fp := [...]int{0, 0, 0, 0}
	// check whether there exists a substring with length k, that after modify it, will make all the string s balance
	check := func(k int) bool {
		for i := 0; i+k <= n; i++ {
			j := i + k

			for u := 0; u < 4; u++ {
				fp[u] = dp[j][u] - dp[i][u]
				fp[u] = dp[n][u] - fp[u]
			}
			// now fp[u] contains element u outside of substring [i...j)
			var v int
			for u := 0; u < 4; u++ {
				if fp[u] > m {
					v = k + 1
					break
				}
				v += m - fp[u]
			}
			// v is count of all the element needed to make s balanced from substring [i....j)
			if v == k {
				return true
			}
		}
		return false
	}

	left, right := 0, n

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
