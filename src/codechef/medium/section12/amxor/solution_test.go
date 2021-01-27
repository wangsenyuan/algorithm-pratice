package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{3, 3}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{3, 3, 3}
	expect := 16
	runSample(t, n, A, expect)
}
func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 4}
	expect := 6
	runSample(t, n, A, expect)
}
