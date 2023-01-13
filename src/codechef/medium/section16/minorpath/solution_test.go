package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1, 1, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 1, 3}
	expect := 3
	runSample(t, A, expect)
}
func TestSample3(t *testing.T) {
	A := []int{1, 0, 2, 3, 3}
	expect := -1
	runSample(t, A, expect)
}
