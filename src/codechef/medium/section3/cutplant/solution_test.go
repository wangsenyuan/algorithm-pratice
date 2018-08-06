package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect int) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{3, 1, 3}
	B := []int{2, 1, 2}
	expect := 2
	runSample(t, n, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	A := []int{1, 3, 4, 5, 1, 2, 3}
	B := []int{1, 2, 1, 2, 1, 1, 1}
	expect := 3
	runSample(t, n, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{2, 3, 9}
	B := []int{2, 3, 9}
	expect := 0
	runSample(t, n, A, B, expect)
}

func TestSample4(t *testing.T) {
	n := 2
	A := []int{1, 2}
	B := []int{2, 1}
	expect := -1
	runSample(t, n, A, B, expect)
}
