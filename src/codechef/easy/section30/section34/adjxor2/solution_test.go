package main

import "testing"

func runSample(t *testing.T, arr []int, x int, expect int) {
	res := int(solve(arr, x))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2}
	x := 1
	expect := 3
	runSample(t, arr, x, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 2, 3, 3}
	x := 1
	expect := 15
	runSample(t, arr, x, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{5, 3, 6, 2, 8}
	x := 2
	expect := 43
	runSample(t, arr, x, expect)
}
