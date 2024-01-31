package main

import "testing"

func runSample(t *testing.T, w []int, b []int, expect int) {
	res := solve(w, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2, 3, 1}
	expect := 12
	runSample(t, a, b, expect)
}
