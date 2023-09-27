package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{0, 1, 1, 0}
	expect := 4
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0}
	expect := 6
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{0, 1, 0, 1, 0, 1}
	expect := 3
	runSample(t, arr, expect)
}
