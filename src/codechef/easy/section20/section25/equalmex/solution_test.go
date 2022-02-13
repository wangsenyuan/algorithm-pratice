package main

import "testing"

func runSample(t *testing.T, arr []int, expect bool) {
	res := solve(arr)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 0, 0, 1}
	expect := false
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 0, 1, 1}
	expect := true
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 3, 2, 3, 3, 2}
	expect := true
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2}
	expect := false
	runSample(t, arr, expect)
}
