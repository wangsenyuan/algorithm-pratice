package main

import "testing"

func runSample(t *testing.T, A []int, k int, expect int64) {
	res := solve(A, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 1, 2, 1, 3, 1}
	k := 3
	var expect int64 = 2
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 5, 5, 5, 5}
	k := 1
	var expect int64 = 16
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 3, 7}
	k := 2
	var expect int64 = 0
	runSample(t, A, k, expect)
}

func TestSample4(t *testing.T) {
	A := []int{4, 2}
	k := 0
	var expect int64 = 6
	runSample(t, A, k, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 10, 10, 1, 10, 2, 7, 10, 3}
	k := 2
	var expect int64 = 16
	runSample(t, A, k, expect)
}
