package main

import "testing"

func runSample(t *testing.T, n int, m int, A []int, x, y int, expect int64) {
	res := solve(n, m, A, x, y)

	if res != expect {
		t.Errorf("sample %d %d %v %d %d, expect %d, but got %d", n, m, A, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	A := []int{1, 2, 3}
	x, y := 0, 1
	runSample(t, n, m, A, x, y, 7)
}
