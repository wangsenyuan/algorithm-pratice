package p2223

func sumScores(s string) int64 {
	S := s + "#" + s
	n := len(S)

	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && S[z[i]] == S[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	var res int64

	for i := len(s) + 1; i < n; i++ {
		res += int64(z[i])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
