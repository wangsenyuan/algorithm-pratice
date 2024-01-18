package main

import "testing"

func runSample(t *testing.T, n int, intervals [][]int, expect int) {
	res := solve(n, intervals)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	expect := 5
	runSample(t, n, nil, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	intervals := [][]int{
		{1, 3},
	}
	expect := 0
	runSample(t, n, intervals, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	intervals := [][]int{
		{3, 4},
		{6, 9},
	}
	expect := 4
	runSample(t, n, intervals, expect)
}

func TestSample4(t *testing.T) {
	n := 1000
	intervals := [][]int{
		{100, 701},
		{200, 801},
		{300, 901},
	}
	expect := 839415253
	runSample(t, n, intervals, expect)
}
