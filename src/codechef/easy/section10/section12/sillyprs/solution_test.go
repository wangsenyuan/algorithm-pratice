package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{4, 5, 6}
	B := []int{1, 2, 3}
	expect := 10
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{4, 8, 6, 4, 1}
	B := []int{2, 5, 7, 4, 7}
	expect := 23
	runSample(t, n, A, B, expect)
}
