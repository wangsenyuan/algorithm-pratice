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
	A := []int{2, 3, 2}
	B := []int{10, 3, 4}
	expect := 17
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{7, 5, 3, 4}
	B := []int{2, 3, 1, 3}
	expect := 19
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{10, 1}
	B := []int{1, 10}
	expect := 20
	runSample(t, n, A, B, expect)
}