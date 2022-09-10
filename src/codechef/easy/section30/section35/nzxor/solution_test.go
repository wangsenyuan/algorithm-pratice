package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 4}
	expect := 2
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 0, 0}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	expect := 0
	runSample(t, arr, expect)
}
