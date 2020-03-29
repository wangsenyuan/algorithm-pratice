package main

import "testing"

func runSample(t *testing.T, n int) {
	a, b := solve(n)

	g := gcd(a, b)
	l := a * b / g

	if g+l != n {
		t.Errorf("Sample %d, result %d %d, not correct", n, a, b)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	for i := 2; i < 100; i++ {
		runSample(t, i)
	}
}
