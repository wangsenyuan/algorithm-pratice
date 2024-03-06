package main

import "testing"

func runSample(t *testing.T, n int, days [][]int, expect int) {
	res := solve(n, days)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	days := [][]int{
		{1, 2},
		{1, 2},
		{1, 1},
	}
	expect := 1
	runSample(t, n, days, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	days := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
	}
	expect := 2
	runSample(t, n, days, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	days := [][]int{
		{1, 5},
		{6, 10},
		{2, 2},
		{3, 7},
		{5, 8},
		{1, 4},
	}
	expect := 3
	runSample(t, n, days, expect)
}

func TestSample4(t *testing.T) {
	n := 100
	days := [][]int{
		{1, 100},
		{1, 100},
		{1, 100},
		{1, 100},
		{1, 100},
		{1, 100},
	}
	expect := 0
	runSample(t, n, days, expect)
}

func TestSample5(t *testing.T) {
	n := 1000
	days := [][]int{
		{1, 1},
		{1, 1},
	}
	expect := 1000
	runSample(t, n, days, expect)
}

func TestSample6(t *testing.T) {
	n := 20
	days := [][]int{
		{9, 20},
		{3, 3},
		{10, 11},
		{11, 13},
		{6, 18},
	}
	expect := 15
	runSample(t, n, days, expect)
}

func TestSample7(t *testing.T) {
	n := 20
	days := [][]int{
		{3, 14},
		{9, 19},
		{5, 18},
		{3, 9},
		{1, 16},
		{10, 18},
		{16, 17},
		{6, 15},
		{3, 9},
		{13, 17},
	}
	expect := 4
	runSample(t, n, days, expect)
}
