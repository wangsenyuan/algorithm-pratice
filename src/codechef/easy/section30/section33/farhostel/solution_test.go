package main

import "testing"

func runSample(t *testing.T, A [][]int, B [][]int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{{6, 7, 0}}
	B := [][]int{{9, 1, 0}}
	var expect int64 = 13
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{89, 0},
		{67, 99},
	}
	B := [][]int{
		{12, 44},
		{23, 54},
	}
	var expect int64 = 299
	runSample(t, A, B, expect)
}
