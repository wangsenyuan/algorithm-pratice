package main

import "testing"

func runSample(t *testing.T, n int, G [][]int) {
	res := solve(n, G)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			tmp := max(max(res[i][j], max(res[i+1][j], res[i][j+1])), res[i+1][j])
			if tmp != G[i][j] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n := 3
	G := [][]int{
		{7, 7},
		{9, 8},
	}
	runSample(t, n, G)
}
