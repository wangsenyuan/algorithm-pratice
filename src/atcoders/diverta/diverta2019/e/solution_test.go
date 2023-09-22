package main

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := solve(arr)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 2, 2}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 147483634
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{1, 2, 5, 3, 3, 6, 1, 1, 8, 8, 0, 3, 3, 4, 6, 6, 4, 0, 7, 2, 5, 4, 6, 2}
	expect := 292
	runSample(t, arr, expect)
}
