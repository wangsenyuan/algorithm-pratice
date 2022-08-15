package main

import "testing"

func runSample(t *testing.T, A [][]int, expect string) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{1, 0, 0},
	}
	expect := "010"
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{0, 1, 1, 0},
		{0, 0, 2, 0},
		{0, 2, 0, 0},
		{1, 1, 1, 0},
	}
	expect := "0001"
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{0, 2, 2, 1, 0},
		{2, 0, 1, 1, 0},
		{2, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 1, 0, 0},
	}
	expect := "11001"
	runSample(t, A, expect)
}
