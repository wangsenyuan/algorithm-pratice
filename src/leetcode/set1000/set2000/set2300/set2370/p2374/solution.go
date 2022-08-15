package p2374

func edgeScore(edges []int) int {
	n := len(edges)
	score := make([]int, n)

	for i := 0; i < n; i++ {
		score[edges[i]] += i
	}

	best := 0

	for i := 0; i < n; i++ {
		if score[i] > score[best] {
			best = i
		}
	}

	return best
}
