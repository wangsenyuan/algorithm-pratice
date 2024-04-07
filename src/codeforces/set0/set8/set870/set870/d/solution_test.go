package main

import "testing"

func runSample(t *testing.T, p []int) {
	n := len(p)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[p[i]] = i
	}
	ask := func(i, j int) int {
		return p[i] ^ b[j]
	}

	_, ans := solve(n, ask)
	sna := make([]int, n)

	for i := 0; i < n; i++ {
		sna[ans[i]] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if p[i]^b[j] != ans[i]^sna[j] {
				t.Fatalf("Sample result %v, not correct", ans)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	p := []int{0, 1, 2}
	runSample(t, p)
}
