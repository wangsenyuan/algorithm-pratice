package main

import "testing"

func runSample(t *testing.T, figures [][]int, P []int64, expect int64) {
	res := solve(figures, P)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	figures := [][]int{
		{5, 10},
	}
	P := []int64{3, 6}
	var expect int64 = 70
	runSample(t, figures, P, expect)
}

func TestSample2(t *testing.T) {
	figures := [][]int{
		{3, 8},
		{5, 10},
	}
	P := []int64{20}
	var expect int64 = 74
	runSample(t, figures, P, expect)
}

func TestSample3(t *testing.T) {
	figures := [][]int{
		{10, 3},
		{20, 2},
		{30, 1},
	}
	P := []int64{30, 50, 60}
	var expect int64 = 200
	runSample(t, figures, P, expect)
}

func TestSample4(t *testing.T) {
	figures := [][]int{
		{5, 9},
		{63, 3},
		{30, 4},
		{25, 6},
		{48, 2},
		{29, 9},
	}
	P := []int64{105, 137, 172, 192, 632, 722, 972, 981}
	var expect int64 = 2251
	runSample(t, figures, P, expect)
}
