package main

import "testing"

func runSample(t *testing.T, n int, C []int, A, B []int, expect int64) {
	res := solve(n, C, A, B)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	C := []int{3, 4, 3, 3}
	A := []int{-1, 1, 2, 2}
	B := []int{-1, 2, 2, 3}
	var expect int64 = 7
	runSample(t, n, C, A, B, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	C := []int{5, 6}
	A := []int{-1, 5}
	B := []int{-1, 1}
	var expect int64 = 11
	runSample(t, n, C, A, B, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	C := []int{3, 5 ,2}
	A := []int{-1, 1, 1}
	B := []int{-1, 3, 5}
	var expect int64 = 8
	runSample(t, n, C, A, B, expect)
}
