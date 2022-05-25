package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 7, 4, 4, 2, 9, 2, 1, 9, 3}
	expect := 15
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 1, 3, 5}
	expect := 6
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 1}
	expect := 4
	runSample(t, A, expect)
}
