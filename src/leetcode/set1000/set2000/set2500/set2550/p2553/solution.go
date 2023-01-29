package p2553

import "sort"

func putMarbles(weights []int, k int) int64 {
	// k <= 1e5
	n := len(weights)
	// n <= 1e5
	// w[0] 肯定放入第一个背包
	var sum int64

	for _, w := range weights {
		sum += int64(w)
	}

	// k * n 太多，无法dp
	// 假设在某个地方分开，那么score += w[i] + w[i+1]
	P := make([]Pair, n-1)
	for i := 0; i+1 < n; i++ {
		P[i] = Pair{int64(weights[i]) + int64(weights[i+1]), i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].first < P[j].first
	})

	max := int64(weights[0]) + int64(weights[n-1])
	min := max
	for i := 0; i+1 < k; i++ {
		j := len(P) - 1 - i
		max += P[j].first
		min += P[i].first
	}
	return max - min
}

type Pair struct {
	first  int64
	second int
}
