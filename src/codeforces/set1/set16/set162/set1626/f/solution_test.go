package main

import "testing"

func runSample(t *testing.T, n int, a0 int, x int, y int, m int, k int, expect int) {
	res := solve(n, a0, x, y, m, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a0, x, y, k, m := 3, 10, 3, 5, 13, 88
	expect := 382842030
	runSample(t, n, a0, x, y, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, a0, x, y, k, m := 2, 15363, 270880, 34698, 17, 2357023
	expect := 319392398
	runSample(t, n, a0, x, y, m, k, expect)
}
