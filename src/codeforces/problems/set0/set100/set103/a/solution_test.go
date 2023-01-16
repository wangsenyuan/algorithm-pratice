package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1}
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2}
	expect := 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{10}
	expect := 10
	runSample(t, A, expect)
}
