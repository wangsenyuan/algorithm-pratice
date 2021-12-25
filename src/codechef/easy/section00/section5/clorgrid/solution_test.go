package main

import "testing"

func runSample(t *testing.T, n, m int, G []string, expect bool) {
	res := solve(n, m, G)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", n, m, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	G := []string{
		"..#",
		"...",
		"...",
	}
	runSample(t, n, m, G, true)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	G := []string{
		"...",
		".#.",
		"...",
	}
	runSample(t, n, m, G, false)
}
