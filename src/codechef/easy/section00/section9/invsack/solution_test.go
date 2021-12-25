package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, W int, w []int, v []int) {
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)

		for j := 0; j <= W; j++ {
			dp[i][j] = -1
		}
	}

	for j := 0; j <= W; j++ {
		dp[0][j] = 0
	}

	for i := 0; i <= N; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= N; i++ {
		for j := W; j >= 0; j-- {
			dp[i][j] = max(dp[i-1][j], dp[i][j])

			if j-w[i-1] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}

	query := func(i, j int) int {
		if i > N {
			i = N
		}
		if j > W {
			j = W
		}
		return dp[i][j]
	}

	nw := make([]int, N)
	nv := make([]int, N)

	solve(N, query, nw, nv)

	if !reflect.DeepEqual(w, nw) && !reflect.DeepEqual(v, nv) {
		t.Errorf("Sample expect %v %v, but got %v %v", w, v, nw, nv)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	w := []int{3, 7}
	v := []int{2, 3}
	N := 2
	runSample(t, N, 10, w, v)
}
