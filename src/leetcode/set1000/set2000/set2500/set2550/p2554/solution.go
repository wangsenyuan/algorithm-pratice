package p2554

func minimumScore(s string, t string) int {
	n := len(t)
	R := make([]int, n)
	for i := 0; i < n; i++ {
		R[i] = -1
	}
	best := n

	for i, j := n-1, len(s)-1; i >= 0; i-- {
		for j >= 0 && s[j] != t[i] {
			j--
		}
		if j < 0 {
			break
		}
		best = i
		R[i] = j
		j--
	}

	for i, j, k := 0, 0, 0; i < n; i++ {
		for j < len(s) && s[j] != t[i] {
			j++
		}
		if j == len(s) {
			break
		}
		if k <= i {
			k = i + 1
		}
		for k < n && R[k] <= j {
			k++
		}
		if k == n || R[k] > j {
			best = min(best, k-i-1)
		}
		j++
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
