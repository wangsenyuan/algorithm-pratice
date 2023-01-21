package main

import "testing"

func runSample(t *testing.T, n int, quarries [][]int64, expect string) {
	res := solve(n, quarries)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	quarries := [][]int64{
		{2, 1},
		{3, 2},
	}
	expect := tolik

	runSample(t, n, quarries, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	quarries := [][]int64{
		{1, 1},
		{1, 1},
		{1, 1},
		{1, 1},
	}
	expect := bolik

	runSample(t, n, quarries, expect)
}
