package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := int(solve(arr))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 1, 2, 3, 8, 2, 1, 4}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 2, 1, 6, 3, 4, 1, 2, 8, 3}
	expect := 10
	runSample(t, arr, expect)
}
