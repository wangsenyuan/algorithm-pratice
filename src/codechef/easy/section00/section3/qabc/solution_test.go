package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect bool) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %t, but got %t", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{0, 0, 0, 0, 0}
	B := []int{1, 2, 4, 2, 3}
	expect := true
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{0, 0, 0, 0, 0}
	B := []int{1, 2, 4, 2, 4}
	expect := false
	runSample(t, n, A, B, expect)
}
