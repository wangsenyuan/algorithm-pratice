package main

import "testing"

func runSample(t *testing.T, n int, m int, C [][]int, k int, op2 int64) {
	rk, rop2 := solve(n, m, C)

	if k != rk || rop2 != op2 {
		t.Errorf("Sample expect %d %d, but got %d %d", k, op2, rk, rop2)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 9
	C := [][]int{
		{0, 1, 2, 0, 0, 2, 1, 1, 0},
		{0, 1, 1, 1, 2, 0, 0, 2, 0},
		{0, 1, 2, 0, 0, 1, 2, 1, 0},
	}
	k := 3
	var op2 int64 = 1
	runSample(t, n, m, C, k, op2)
}

func TestSample2(t *testing.T) {
	n, m := 3, 7
	C := [][]int{
		{25, 15, 20, 15, 25, 20, 20},
		{26, 14, 20, 14, 26, 20, 20},
		{25, 15, 20, 15, 20, 20, 25},
	}
	k := 3
	var op2 int64 = 10
	runSample(t, n, m, C, k, op2)
}
