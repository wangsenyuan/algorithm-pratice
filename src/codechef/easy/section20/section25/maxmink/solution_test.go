package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, B []int, expect int) {
	res := solve(n, k, A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	A := []int{4, 2, 3, 1, 4}
	B := []int{3, 2, 5, 5, 1}
	expect := 9
	runSample(t, n, k, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	A := []int{1, 2, 3, 4}
	B := []int{4, 3, 2, 1}
	expect := 5
	runSample(t, n, k, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, k := 6, 3
	A := []int{8, 10, 3, 6, 7, 2}
	B := []int{4, 8, 4, 1, 1, 6}
	expect := 18
	runSample(t, n, k, A, B, expect)
}
