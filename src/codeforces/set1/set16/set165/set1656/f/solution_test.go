package main

import "testing"

func runSample(t *testing.T, A []int, ok bool, expect int64) {
	ok1, res := solve(A)

	if ok != ok1 {
		t.Fatalf("Sample expect %t, but got %t", ok, ok1)
	}
	if ok && expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 0}
	ok := false
	var expect int64
	runSample(t, A, ok, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-1, 1}
	ok := true
	var expect int64 = -1
	runSample(t, A, ok, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, -4}
	ok := true
	var expect int64 = -18
	runSample(t, A, ok, expect)
}
