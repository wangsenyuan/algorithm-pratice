package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := int(solve(A, B))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	B := []int{3, 2, 1}
	expect := 1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4}
	B := []int{5, 6, 6, 7}
	expect := 0
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{10, 10, 10}
	B := []int{10, 10, 10}
	expect := 3
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{4, 8, 1, 1}
	B := []int{8, 4, 1, 1}
	expect := 2
	runSample(t, A, B, expect)
}
