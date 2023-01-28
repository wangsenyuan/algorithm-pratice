package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr[0], arr[1], arr[2], arr[3])

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{3, 6, 1, 4}
	expect := 2

	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 1, 4, 4}
	expect := 0

	runSample(t, arr, expect)
}
