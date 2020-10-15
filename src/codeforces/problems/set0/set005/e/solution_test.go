package main

import "testing"

func runSample(t *testing.T, n int, arr []int, expect int64) {
	res := solve(n, arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	arr := []int{1, 2, 4, 5, 3}
	var expect int64 = 7
	runSample(t, n, arr, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	arr := []int{2, 2, 2}
	var expect int64 = 3
	runSample(t, n, arr, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	arr := []int{2, 1, 2}
	var expect int64 = 3
	runSample(t, n, arr, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	arr := []int{1, 2, 3}
	var expect int64 = 3
	runSample(t, n, arr, expect)
}
