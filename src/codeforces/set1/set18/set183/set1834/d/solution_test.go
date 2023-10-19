package main

import "testing"

func runSample(t *testing.T, segs [][]int, m int, expect int) {
	res := solve(segs, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	segs := [][]int{
		{2, 6},
		{4, 8},
		{2, 7},
		{1, 5},
	}
	m := 8
	expect := 6
	runSample(t, segs, m, expect)
}

func TestSample2(t *testing.T) {
	segs := [][]int{
		{1, 3},
		{2, 3},
		{2, 2},
	}
	m := 3
	expect := 4
	runSample(t, segs, m, expect)
}
func TestSample3(t *testing.T) {
	segs := [][]int{
		{1, 5},
		{1, 5},
		{1, 5},
	}
	m := 5
	expect := 0
	runSample(t, segs, m, expect)
}

func TestSample4(t *testing.T) {
	segs := [][]int{
		{1, 1},
		{3, 3},
		{5, 5},
	}
	m := 5
	expect := 2
	runSample(t, segs, m, expect)
}

func TestSample5(t *testing.T) {
	segs := [][]int{
		{1, 7},
		{1, 3},
		{3, 3},
		{4, 5},
	}
	m := 7
	expect := 12
	runSample(t, segs, m, expect)
}

func TestSample6(t *testing.T) {
	segs := [][]int{
		{1, 3},
		{2, 4},
	}
	m := 4
	expect := 2
	runSample(t, segs, m, expect)
}
