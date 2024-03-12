package main

import "testing"

func runSample(t *testing.T, shows [][]int, expect bool) {
	res := solve(shows)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	shows := [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
	}
	expect := true
	runSample(t, shows, expect)
}

func TestSample2(t *testing.T) {
	shows := [][]int{
		{1, 2},
		{2, 3},
		{2, 3},
		{1, 2},
	}
	expect := false
	runSample(t, shows, expect)
}

func TestSample3(t *testing.T) {
	shows := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
	}
	expect := true
	runSample(t, shows, expect)
}

func TestSample4(t *testing.T) {
	shows := [][]int{
		{1, 3},
		{1, 4},
		{4, 10},
		{5, 8},
		{9, 11},
	}
	expect := true
	runSample(t, shows, expect)
}
