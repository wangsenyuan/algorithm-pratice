package main

import "testing"

func runSample(t *testing.T, n, m int, A, B [][]int, expect int) {
	res := solve(n, m, A, B)
	if int(res) != expect {
		t.Errorf("sample %d %d %v %v, should give %d, but got %d", n, m, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{1, 1},
		{1, 1},
	}
	B := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := 3
	runSample(t, n, m, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{1, 9},
		{9, 1},
	}
	B := [][]int{
		{9, 1},
		{1, 9},
	}
	expect := -1
	runSample(t, n, m, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, m := 1, 4
	A := [][]int{
		{4, 5, 7, 1},
	}
	B := [][]int{
		{2, 3, 4, 5},
	}
	expect := 9
	runSample(t, n, m, A, B, expect)
}
