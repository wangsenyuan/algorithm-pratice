package main

import "testing"

func runSample(t *testing.T, n, k, p int, arr []int, expect int) {
	res := solve(n, k, p, arr)
	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d but got %d", n, k, p, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, p := 4, 3, 0
	arr := []int{4, 1, 9, 5}
	expect := 9
	runSample(t, n, k, p, arr, expect)
}

func TestSample2(t *testing.T) {
	n, k, p := 4, 3, 1
	arr := []int{4, 1, 9, 5}
	expect := 1
	runSample(t, n, k, p, arr, expect)
}
