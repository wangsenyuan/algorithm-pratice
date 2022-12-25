package p2515

func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)

	best := -1

	for i := 0; i < n; i++ {
		j := (startIndex + i) % n
		if words[j] == target {
			x := min(i, (n - i))
			if best < 0 || best > x {
				best = x
			}
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
