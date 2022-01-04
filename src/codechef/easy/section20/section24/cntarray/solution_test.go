package main

import "testing"

func runSample(t *testing.T, n int, m int, B []int, expect int) {
	res := solve(n, m, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 4
	B := []int{5, 1, 2}
	expect := 14
	runSample(t, n, m, B, expect)
}
