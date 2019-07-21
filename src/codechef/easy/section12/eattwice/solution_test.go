package main

import "testing"

func runSample(t *testing.T, n, m int, D []int, V []int, expect int) {
	res := solve(n, m, D, V)
	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, D, V, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 6
	D := []int{5, 1, 2}
	V := []int{7, 9, 5}
	expect := 16
	runSample(t, n, m, D, V, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 7
	D := []int{5, 2, 5}
	V := []int{8, 5, 10}
	expect := 15
	runSample(t, n, m, D, V, expect)
}
