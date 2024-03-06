package main

import "testing"

func runSample(t *testing.T, n int, w []int, expect int) {
	res := solve(n, w)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 12
	w := []int{1, 3, 3, 7, 8, 5, 2, 2, 2, 2, 4, 2}
	expect := 36
	runSample(t, n, w, expect)
}
