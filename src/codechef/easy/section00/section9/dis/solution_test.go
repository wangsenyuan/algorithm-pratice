package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect bool) {
	res := solve(n, A, B)
	if res != expect {
		t.Errorf("Sample %d %v %v, expect %t, but got %t", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{0, 1, 1}
	B := []int{1, 1, 0}
	expect := true
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{0, 1, 5}
	B := []int{5, 1, 0}
	expect := false
	runSample(t, n, A, B, expect)
}
