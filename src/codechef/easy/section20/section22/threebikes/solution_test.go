package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 5, 6}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 4}
	expect := 0
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 6, 8, 9, 10, 13}
	expect := 24
	runSample(t, A, expect)
}
