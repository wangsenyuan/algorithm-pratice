package main

import "testing"

func runSample(t *testing.T, m int, segs [][]int, expect int) {
	res := solve(m, segs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	segs := [][]int{
		{2, 2},
	}
	expect := 1
	runSample(t, m, segs, expect)
}
