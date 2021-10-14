package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := int(solve(A, B))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	B := []int{2, 1}
	expect := -1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 2, 1}
	B := []int{1, 2, 3}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3}
	B := []int{1, 2, 3}
	expect := 0
	runSample(t, A, B, expect)
}
