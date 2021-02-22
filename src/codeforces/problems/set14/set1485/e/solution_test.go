package main

import "testing"

func runSample(t *testing.T, n int, P []int, A []int, expect int64) {
	res := solve(n, P, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 14
	P := []int{1, 1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 8, 8}
	A := []int{2, 3, 7, 7, 6, 9, 5, 9, 7, 3, 6, 6, 5}
	var expect int64 = 14
	runSample(t, n, P, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	P := []int{1, 2, 2, 3, 4}
	A := []int{32, 78, 69, 5, 41}
	var expect int64 = 45
	runSample(t, n, P, A, expect)
}

func TestSample3(t *testing.T) {
	n := 15
	P := []int{1, 15, 1, 10, 4, 9, 11, 2, 4, 1, 8, 6, 10, 11}
	A := []int{62, 13, 12, 43, 39, 65, 42, 86, 25, 38, 19, 19, 43, 62}
	var expect int64 = 163
	runSample(t, n, P, A, expect)
}

func TestSample4(t *testing.T) {
	n := 15
	P := []int{11, 2, 7, 6, 9, 8, 10, 1, 1, 1, 5, 3, 15, 2}
	A := []int{50, 19, 30, 35, 9, 45, 13, 24, 8, 44, 16, 26, 10, 40}
	var expect int64 = 123
	runSample(t, n, P, A, expect)
}

func TestSample5(t *testing.T) {
	n := 2
	P := []int{1}
	A := []int{806680823}
	var expect int64 = 0
	runSample(t, n, P, A, expect)
}
