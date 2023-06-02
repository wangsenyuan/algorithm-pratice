package main

import "testing"

func runSample(t *testing.T, x, y []int64, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int64{1, 5, 4, 1}
	y := []int64{1, 2, 5, 4}
	expect := 4
	runSample(t, x, y, expect)
}
