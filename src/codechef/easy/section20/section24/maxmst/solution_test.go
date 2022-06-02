package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, expect int) {
	res := int(solve(n, m, A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 4
	A := []int{1, 2, 3, 4}
	expect := 7
	runSample(t, n, m, A, expect)
}
