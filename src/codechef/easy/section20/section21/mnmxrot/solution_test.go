package main

import "testing"

func runSample(t *testing.T, A []int, S string, expect int) {
	res := solve(len(A), A, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 1, 2, 6}
	S := "BWBW"
	expect := 4
	runSample(t, A, S, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 8 ,2 ,9 ,3}
	S := "BBWBB"
	expect := 8
	runSample(t, A, S, expect)
}
