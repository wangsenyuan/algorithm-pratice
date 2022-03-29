package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2, 3}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	expect := 1
	runSample(t, A, expect)
}
