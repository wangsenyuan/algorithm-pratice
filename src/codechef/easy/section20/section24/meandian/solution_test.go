package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1000000000, 1, 1000000000, 1}
	expect := 500000002
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 3, 3}
	expect := 4
	runSample(t, A, expect)
}
