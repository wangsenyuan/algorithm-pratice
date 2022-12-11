package p2495

import (
	"math"
	"sort"
)

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	adj := make([][]int, n)

	add := func(u int, x int) {
		if len(adj[u]) == 0 {
			adj[u] = make([]int, 0, 1)
		}
		adj[u] = append(adj[u], x)
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		if vals[u] > 0 {
			add(v, vals[u])
		}
		if vals[v] > 0 {
			add(u, vals[v])
		}
	}

	best := math.MinInt32

	for i := 0; i < n; i++ {
		sort.Ints(adj[i])
		sum := vals[i]
		var x int
		for j := len(adj[i]) - 1; j >= 0 && x < k; j-- {
			if adj[i][j] <= 0 {
				break
			}
			sum += adj[i][j]
			x++
		}
		best = max(best, sum)
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
