package main

import "testing"

func runSample(t *testing.T, n int, E [][]int) {
	res := solve(n, E)

	dp := make([][]int, n)
	fp := make([][]uint64, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		fp[i] = make([]uint64, n)
	}

	adj := getGraph(n, E)

	que := make([]int, n)
	bfs := func(x int) {
		var front, end int
		que[end] = x
		end++

		for i := 0; i < n; i++ {
			dp[x][i] = 0
			fp[x][i] = 1
		}

		dp[x][x] = 1
		fp[x][x] = res[x]

		for front < end {
			u := que[front]
			front++
			for _, v := range adj[u] {
				if dp[x][v] == 0 {
					dp[x][v] = dp[x][u] + 1
					fp[x][v] = fp[x][u] * res[v]
					que[end] = v
					end++
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		bfs(i)
		for j := 0; j < n; j++ {
			x := fp[i][j]
			y := uint64(dp[i][j])

			if x%y != 0 {
				t.Fatalf("Sample fail at path %d -> %d", i, j)
			}
		}
	}
}

// func TestSample1(t *testing.T) {
// 	n := 6
// 	E := [][]int{
// 		{1, 2},
// 		{2, 3},
// 		{1, 4},
// 		{5, 1},
// 		{6, 5},
// 	}
// 	runSample(t, n, E)
// }

func TestSample2(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2},
	}
	runSample(t, n, E)
}
