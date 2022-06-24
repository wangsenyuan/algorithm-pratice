package main

import "testing"

func runSample(t *testing.T, A, B, N int64, expect int) {
	res := solve(A, B, N)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var A, B, N int64 = 5, 5, 2
	expect := 0
	runSample(t, A, B, N, expect)
}

func TestSample2(t *testing.T) {
	var A, B, N int64 = 3, 7, 8
	expect := 1
	runSample(t, A, B, N, expect)
}
func TestSample3(t *testing.T) {
	var A, B, N int64 = 8, 11, 1
	expect := -1
	runSample(t, A, B, N, expect)
}