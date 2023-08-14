package main

import "testing"

func runSample(t *testing.T, n int64, k int64, expect int64) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var n, k int64 = 8, 3
	var expect int64 = 4
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	var n, k int64 = 6, 6
	var expect int64 = 3
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	var n, k int64 = 7, 1
	var expect int64 = 6
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	var n, k int64 = 1, 1
	var expect int64 = 0
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	var n, k int64 = 2, 2
	var expect int64 = 1
	runSample(t, n, k, expect)
}
