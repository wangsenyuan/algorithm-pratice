package main

import "testing"

func runSample(t *testing.T, grid []string, n, m, l, k int, expect int64) {
	bs := make([][]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = []byte(grid[i])
	}
	res := solve(n, m, bs, l, k)
	if res != expect {
		t.Errorf("sample: %d %d %d %d %v, should give %d, but got %d", n, m, l, k, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, l, k := 3, 2, 2, 1
	grid := []string{
		"11",
		"01",
		"10",
	}
	runSample(t, grid, n, m, l, k, 2)
}

func TestSample2(t *testing.T) {
	n, m, l, k := 3, 3, 2, 2
	grid := []string{
		"111",
		"101",
		"111",
	}
	runSample(t, grid, n, m, l, k, 5)
}
