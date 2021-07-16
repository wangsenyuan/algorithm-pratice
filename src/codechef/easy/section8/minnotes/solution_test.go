package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := int(solve(n, A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{8, 4, 2}
	expect := 4
	runSample(t, n, A, expect)
}
