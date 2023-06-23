package main

import "testing"

func runSample(t *testing.T, A []int, S string, expect int) {
	res := solve(A, S)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	S := "L"
	expect := 5
	runSample(t, A, S, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3, 4, 5}
	S := "LR"
	expect := 416666673
	runSample(t, A, S, expect)
}
