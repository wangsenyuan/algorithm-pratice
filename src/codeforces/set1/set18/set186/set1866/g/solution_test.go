package main

import "testing"

func runSample(t *testing.T, a []int, d []int, expect int) {
	res := solve(a, d)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{7, 4, 2, 0, 5, 8, 3}
	d := []int{4, 0, 0, 1, 3, 1, 3}
	expect := 5
	runSample(t, a, d, expect)
}
