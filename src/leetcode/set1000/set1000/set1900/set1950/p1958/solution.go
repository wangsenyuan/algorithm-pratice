package p1958

func maxProduct(s string) int64 {
	n := len(s)
	P := make([]int, n)

	C := 0
	R := 0
	for i := 0; i < n; i++ {
		j := 2*C - i // equals to i' = C - (i-C)

		if R > i {
			P[i] = min(R-i, P[j])
		}
		// Attempt to expand palindrome centered at i
		for i+1+P[i] < n && i-1-P[i] >= 0 && s[i+1+P[i]] == s[i-1-P[i]] {
			P[i]++
		}

		// If palindrome centered at i expand past R,
		// adjust center based on expanded palindrome.
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}

	arr := make([]int, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = 2 * n
	}

	set := func(arr []int, p, v int, op func(int, int) int) {
		p += n
		arr[p] = op(arr[p], v)
		for p > 0 {
			arr[p>>1] = op(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(arr []int, l, r int, op func(int, int) int, v int) int {
		l += n
		r += n
		res := v
		for l < r {
			if l&1 == 1 {
				res = op(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = op(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		set(arr, i+P[i], i, min)
		// find min j that j + P[j] >= i
		j := get(arr, i, n, min, n)
		dp[i] = 2*(i-j) + 1
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}

	arr2 := make([]int, 2*n)

	var res int64
	var best int
	for i := n - 1; i > 0; i-- {
		// need to find j that, j - P[j] <= i
		set(arr2, i-P[i], i, max)
		j := get(arr2, 0, i+1, max, i)
		//j >= i
		best = max(best, 2*(j-i)+1)
		res = max2(res, int64(best)*int64(dp[i-1]))
	}

	return res
}

func max2(a, b int64) int64 {
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findPalindrome(t string) []int {
	n := len(t)

	P := make([]int, n)
	C := 0
	R := 0
	for i := 1; i < n-1; i++ {
		j := 2*C - i // equals to i' = C - (i-C)

		if R > i {
			P[i] = min(R-i, P[j])
		}
		// Attempt to expand palindrome centered at i
		for i+1+P[i] < n && i-1-P[i] >= 0 && t[i+1+P[i]] == t[i-1-P[i]] {
			P[i]++
		}

		// If palindrome centered at i expand past R,
		// adjust center based on expanded palindrome.
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}
	return P
}
