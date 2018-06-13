package main

import "testing"

func runSample(t *testing.T, n int, C []int, m int, P []int, meals [][]int, expect int64) {
	res := solve(n, C, m, P, meals)
	if res != expect {
		t.Errorf("sample %d %v %d %v %v, should give %d, but got %d", n, C, m, P, meals, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	C := []int{3, 5, 6}
	P := []int{11, 5, 5}
	meals := [][]int{
		{1, 2, 3},
		{1, 2},
		{1, 3},
	}
	expect := int64(10)
	runSample(t, n, C, m, P, meals, expect)
}
