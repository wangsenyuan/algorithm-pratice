package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 6, 3, 4, 2}
	expect := 11
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	expect := 6
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 10}
	expect := 9
	runSample(t, A, expect)
}
