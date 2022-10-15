package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int) {
	res := int(solve(A, k))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{1, 3, 6}
	expect := 4
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	A := []int{1, 2}
	expect := 1
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	A := []int{1, 2, 5}
	expect := 3
	runSample(t, A, k, expect)
}
