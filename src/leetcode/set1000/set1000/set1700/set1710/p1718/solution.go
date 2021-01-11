package p1718

import "sort"

const N = 500

func checkWays(pairs [][]int) int {
	degree := make([]int, N)
	for _, pair := range pairs {
		u, v := pair[0], pair[1]
		degree[u-1]++
		degree[v-1]++
	}
	ps := make([]Pair, 0, N)
	for i := 0; i < N; i++ {
		if degree[i] > 0 {
			ps = append(ps, Pair{degree[i], i})
		}
	}

	n := len(ps)

	G := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		G[ps[i].second] = make(map[int]int)
	}

	for _, pair := range pairs {
		u, v := pair[0], pair[1]
		G[u-1][v-1]++
		G[v-1][u-1]++
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].first > ps[j].first
	})

	if ps[0].first != n-1 {
		return 0
	}

	var pos int
	for pos < n && ps[pos].first == n-1 {
		u := ps[pos].second
		for v := range G[u] {
			delete(G[v], u)
		}
		pos++
	}
	var ans = 1

	if pos > 1 {
		ans = 2
	}

	for pos < n {
		u := ps[pos].second
		k := len(G[u])

		for v := range G[u] {
			if len(G[v]) > k {
				return 0
			}
			if len(G[v]) == k {
				ans = 2
			}
			delete(G[v], u)

			for w := range G[v] {
				if G[u][w] == 0 {
					return 0
				}
			}
		}

		pos++
	}
	// then after all, there should be one tree left
	return ans
}

type Pair struct {
	first, second int
}
