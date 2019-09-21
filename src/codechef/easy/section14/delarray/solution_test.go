package main

import "testing"

func runSample(t *testing.T, n int, arr []int, expect int) {
	res := solve(n, arr)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	arr := []int{1, 1, 2}
	expect := 4
	runSample(t, n, arr, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	arr := []int{2, 4, 3, 5}
	expect := 7
	runSample(t, n, arr, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	arr := []int{1, 1}
	expect := 2
	runSample(t, n, arr, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	arr := []int{1, 1, 1}
	expect := 2
	runSample(t, n, arr, expect)
}
