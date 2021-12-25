package main

import "testing"

func runSample(t *testing.T, n, k, time int, tm [][]int, pm [][]int, expect int) {
	res := solve(n, k, time, tm, pm)
	if res != expect {
		t.Errorf("sample %d %d %d %v %v, expect %d, but got %d", n, k, time, tm, pm, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, time := 1, 0, 5
	tm := [][]int{
		{1, 2, 3},
	}
	pm := [][]int{
		{5, 3, 6},
	}
	expect := 6
	runSample(t, n, k, time, tm, pm, expect)
}

func TestSample2(t *testing.T) {
	n, k, time := 2, 1, 6
	tm := [][]int{
		{1, 2, 1},
		{2, 3, 3},
	}
	pm := [][]int{
		{5, 3, 3},
		{4, 6, 5},
	}
	expect := 11
	runSample(t, n, k, time, tm, pm, expect)
}
