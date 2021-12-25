package main

import "testing"

func runSample(t *testing.T, N int, A []int, Z1 int, Z2 int, expect int) {
	res := solve(N, A, Z1, Z2)
	if res != expect {
		t.Errorf("Sample %d %v %d %d, expect %d, but got %d", N, A, Z1, Z2, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Z1, Z2 := 2, 6, 4
	A := []int{-4, 10}
	expect := 1
	runSample(t, N, A, Z1, Z2, expect)
}

func TestSample2(t *testing.T) {
	N, Z1, Z2 := 1, 1, -1
	A := []int{2}
	expect := 0
	runSample(t, N, A, Z1, Z2, expect)
}

func TestSample3(t *testing.T) {
	N, Z1, Z2 := 2, 0, 7
	A := []int{3, 4}
	expect := 2
	runSample(t, N, A, Z1, Z2, expect)
}
