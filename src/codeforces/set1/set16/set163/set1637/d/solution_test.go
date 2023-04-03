package main

import "testing"

func runSample(t *testing.T, A, B []int, expect int64) {
	res := solve(A, B)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 6, 6, 6}
	B := []int{2, 7, 4, 1}
	var expect int64 = 987
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 7, 2, 4}
	B := []int{2, 5, 3, 5}
	var expect int64 = 914
	runSample(t, A, B, expect)
}
