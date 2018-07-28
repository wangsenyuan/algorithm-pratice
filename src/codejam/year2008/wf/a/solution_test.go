package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, C []int, expect int) {
	res := solve(n, A, B, C)

	if res != expect {
		t.Errorf("Sample %d %v %v %v, expect %d, but got %d", n, A, B, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{10000, 0, 0}
	B := []int{0, 10000, 0}
	C := []int{0, 0, 10000}
	expect := 1
	runSample(t, n, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{5000, 0, 0}
	B := []int{0, 2000, 0}
	C := []int{0, 0, 4000}
	expect := 2
	runSample(t, n, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{0, 3000, 1000, 2000, 1000}
	B := []int{1250, 0, 1000, 1000, 3000}
	C := []int{0, 3000, 1000, 2000, 2000}
	expect := 5
	runSample(t, n, A, B, C, expect)
}
