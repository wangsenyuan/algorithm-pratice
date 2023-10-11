package main

import "testing"

func runSample(t *testing.T, p []int) {
	ask := func(i int, j int) int {
		return gcd(p[i-1], p[j-1])
	}

	res := solve(len(p), ask)

	i, j := res[0], res[1]

	if p[i-1] != 0 && p[j-1] != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	p := []int{1, 0}
	runSample(t, p)
}

func TestSample2(t *testing.T) {
	p := []int{2, 4, 0, 1, 3}
	runSample(t, p)
}
