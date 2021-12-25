package main

import "testing"

func runSample(t *testing.T, n int, x int, c []int, T []int, expect int64) {
	res := solve(n, x, c, T)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 5, 6
	c := []int{8, 3, 1, 5, 0}
	T := []int{1, 1, 1, 1}
	var expect int64 = 3
	runSample(t, n, x, c, T, expect)
}

func TestSample2(t *testing.T) {
	n, x := 5, 1
	c := []int{2, 8, 0, 0, 1}
	T := []int{7, 3, 1, 2}
	var expect int64 = 16
	runSample(t, n, x, c, T, expect)
}
