package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr[0], arr[1], arr[2], arr[3], arr[4], arr[5])

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{4, 3, 1, 2, 2, 191}
	expect := 1228
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{10, 5, 1, 113, 157, 999991}
	expect := 328836201
	runSample(t, arr, expect)
}
