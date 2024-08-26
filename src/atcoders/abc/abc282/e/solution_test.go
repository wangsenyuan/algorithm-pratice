package main

import "testing"

func runSample(t *testing.T, m int, a []int, expect int) {
	res := solve(m, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 100
	a := []int{29, 31, 68, 20, 83, 66, 23, 84, 69, 96, 41, 61, 83, 37, 52, 71, 18, 55, 40, 8}
	expect := 1733
	runSample(t, m, a, expect)
}
