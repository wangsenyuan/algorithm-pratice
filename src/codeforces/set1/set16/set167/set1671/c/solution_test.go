package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int64) {
	res := solve(A, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 2}
	x := 7
	var expect int64 = 11
	runSample(t, A, x, expect)
}

func TestSample2(t *testing.T) {
	A := []int{10, 20, 30, 40, 50}
	x := 9
	var expect int64 = 0
	runSample(t, A, x, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	x := 1
	var expect int64 = 1
	runSample(t, A, x, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 1}
	x := 1000
	var expect int64 = 1500
	runSample(t, A, x, expect)
}
