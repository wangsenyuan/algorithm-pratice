package p5551

func minimumDeletions(s string) int {
	// make sure a precede b
	n := len(s)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			cnt[i]++
		}
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
	}

	sum := cnt[n-1]

	best := min(sum, n-sum)

	for i := 0; i < n; i++ {
		a := cnt[i]
		b := (i + 1) - a
		aa := sum - a
		// make s[0...i] a and s[i+1:] as b
		tmp := b + aa
		best = min(best, tmp)
	}
	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
