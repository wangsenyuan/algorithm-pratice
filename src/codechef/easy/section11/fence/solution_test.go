package main

import "testing"

func runSample(t *testing.T, n, m, k int, plants [][]int, expect int) {
	res := solve(n, m, k, plants)
	if res != expect {
		t.Errorf("sample %d %d %d %v, expect %d, but got %d", n, m, k, plants, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 4, 4, 9
	plants := [][]int{
		{1, 4},
		{2, 1},
		{2, 2},
		{2, 3},
		{3, 1},
		{3, 3},
		{4, 1},
		{4, 2},
		{4, 3},
	}
	expect := 20
	runSample(t, n, m, k, plants, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 4, 4, 1
	plants := [][]int{
		{1, 1},
	}
	expect := 4
	runSample(t, n, m, k, plants, expect)
}
