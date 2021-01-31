package p1744

func checkPartitioning(s string) bool {
	n := len(s)
	if n < 3 {
		return false
	}

	// lets find whether s[0...i] is palindrome or not
	x := check(s)
	y := check(reverse(s))

	for i := 1; i < n-1; i++ {
		// take s[i] as center
		for l := 0; i-l > 0 && i+l < n-1 && s[i-l] == s[i+l]; l++ {
			if x[i-l-1] == 0 && y[n-1-(i+l+1)] == 0 {
				return true
			}
		}
		// take s[i] s[i+1] as center
		for l := 0; i-l > 0 && i+1+l < n-1 && s[i-l] == s[i+1+l]; l++ {
			if x[i-l-1] == 0 && y[n-1-(i+l+2)] == 0 {
				return true
			}
		}
	}

	return false
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func check(s string) []int {
	t := preProcess(s)
	n := len(t)
	P := make([]int, n, n)
	C := 0
	R := 0
	for i := 1; i < n-1; i++ {
		j := 2*C - i // equals to i' = C - (i-C)

		if R > i {
			P[i] = min(R-i, P[j])
		}
		// Attempt to expand palindrome centered at i
		for t[i+1+P[i]] == t[i-1-P[i]] {
			P[i]++
		}
		// If palindrome centered at i expand past R,
		// adjust center based on expanded palindrome.
		if i+P[i] > R {
			C = i
			R = i + P[i]
		}
	}
	res := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		res[i] = i
	}

	for i := 2; i < n-1; i++ {
		a := (i - 1 - P[i]) / 2
		b := a + P[i]
		res[b-1] = min(res[b-1], a)
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func preProcess(s string) string {
	rs := make([]rune, len(s)*2+3, len(s)*2+3)
	i := 0
	rs[i] = '^'

	for _, r := range s {
		i++
		rs[i] = '#'
		i++
		rs[i] = r
	}
	i++
	rs[i] = '#'
	i++
	rs[i] = '$'

	return string(rs)
}
