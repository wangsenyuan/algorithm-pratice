package main

import "testing"

func runSample(t *testing.T, n int, m int, R []int, C [][]int, expect int) {
	res := solve(n, m, R, C)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, R, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	R := []int{2500, 2500, 2520}
	C := [][]int{
		{10, -5, -20},
		{10, 15, 20},
		{-15, 17, 13},
	}
	expect := 2
	runSample(t, n, m, R, C, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 3
	R := []int{2125, 2098}
	C := [][]int{
		{-20, 10, -10},
		{10, 10, -20},
	}
	expect := 2
	runSample(t, n, m, R, C, expect)
}
