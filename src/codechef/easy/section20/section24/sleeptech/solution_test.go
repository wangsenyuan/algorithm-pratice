package main

import "testing"

func runSample(t *testing.T, n, a, b int, S [][]int64, expect int) {
	res := solve(n, a, b, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 3, 1, 5
	S := [][]int64{
		{1, 5},
		{3, 9},
		{6, 9},
	}
	expect := 2
	runSample(t, n, a, b, S, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 3, 7, 54
	S := [][]int64{
		{54, 228},
		{228, 1337},
		{69, 213},
	}
	expect := 2
	runSample(t, n, a, b, S, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 3, 1, 9
	S := [][]int64{
		{100, 100},
		{99, 99},
		{98, 98},
	}
	expect := 0
	runSample(t, n, a, b, S, expect)
}
