package main

import "testing"

func runSample(t *testing.T, n int, P, W, V []int, expect int64) {
	res := solve(n, P, W, V)

	if res != expect {
		t.Errorf("Sample %d %v %v %v, expect %d, but got %d", n, P, W, V, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	W := []int{3, 7}
	P := []int{7, 10}
	V := []int{4, 5}
	expect := int64(9)
	runSample(t, n, P, W, V, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	W := []int{3, 6}
	P := []int{7, 10}
	V := []int{4, 5}
	expect := int64(5)
	runSample(t, n, P, W, V, expect)
}
