package main

import "testing"

func runSample(t *testing.T, A, B [][]int, expect int) {
	res := solve(A, B)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 2},
		{2, 3},
		{3, 1},
	}
	B := [][]int{
		{1, 4},
		{2, 1},
		{3, 3},
	}
	expect := 4
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{19, 79},
		{33, 463},
		{19, 178},
		{33, 280},
	}
	B := [][]int{
		{19, 255},
		{33, 92},
		{34, 25},
		{19, 96},
		{12, 11},
		{19, 490},
		{33, 31},
	}
	expect := 380
	runSample(t, A, B, expect)
}
