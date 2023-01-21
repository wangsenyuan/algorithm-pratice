package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect bool) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %t, but got %t", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{10, 20, 20, 30}
	B := []int{0, 1, 0, 1}
	expect := true
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{5, 15, 4}
	B := []int{0, 0, 0}
	expect := false
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{3, 1, 2}
	B := []int{0, 1, 1}
	expect := true
	runSample(t, n, A, B, expect)
}
