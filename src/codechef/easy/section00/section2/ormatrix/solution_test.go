package main

import (
	"testing"
)

func runSample(t *testing.T, N, M int, A []string, expect [][]int) {
	B := make([][]byte, N)
	for i := 0; i < N; i++ {
		B[i] = []byte(A[i])
	}
	res := solve(N, M, B)
	if !eq(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", N, M, A, expect, res)
	}
}

func eq(a, b [][]int) bool {
	n := len(b)
	m := len(b[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	N, M := 3, 3
	A := []string{
		"010",
		"000",
		"001",
	}
	expect := [][]int{
		{1, 0, 1},
		{2, 1, 1},
		{1, 1, 0},
	}
	runSample(t, N, M, A, expect)
}
