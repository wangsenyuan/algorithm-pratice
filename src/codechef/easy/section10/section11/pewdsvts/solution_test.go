package main

import "testing"

func runSample(t *testing.T, N, A, B, X, Y, Z int, C []int, expect int) {
	res := solve(N, A, B, X, Y, Z, C)
	if res != expect {
		t.Errorf("Sample %d %d %d %d %d %d %v, expect %d, but got %d", N, A, B, X, Y, Z, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{12, 15, 18}
	runSample(t, 3, 10, 15, 5, 10, 100, C, 4)
}

func TestSample2(t *testing.T) {
	C := []int{5, 5, 10}
	runSample(t, 3, 10, 15, 5, 10, 100, C, -1)
}


func TestSample3(t *testing.T) {
	C := []int{100, 100, 100, 100}
	runSample(t, 4, 40, 80, 30, 30, 100, C, 1)
}
