package p2379

func minimumRecolors(blocks string, k int) int {
	var cnt int
	best := len(blocks)
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			cnt++
		}
		if i-k >= 0 {
			if blocks[i-k] == 'W' {
				cnt--
			}
		}
		if i >= k-1 {
			best = min(best, cnt)
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
