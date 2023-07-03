package main

import "testing"

func runSample(t *testing.T, A []int, P []int, expect int64) {
	res := solve(A, P)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 2}
	P := []int{1, 1}
	var expect int64 = 3
	runSample(t, A, P, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1, 3}
	P := []int{1, 1}
	var expect int64 = 4
	runSample(t, A, P, expect)
}
