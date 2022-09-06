package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{3, 1}
	expect := 2
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 3, 3}
	expect := 5
	runSample(t, arr, expect)
}
