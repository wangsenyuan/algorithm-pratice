package main

import "testing"

func runSample(t *testing.T, arr []int, x int, expect int) {
	res := solve(arr, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 0, 0}
	x := 1
	expect := 0
	runSample(t, arr, x, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{1, 1, 0}
	x := 1
	expect := 1
	runSample(t, arr, x, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{0, 1, 0, 1, 1, 1, 0, 0, 1}
	x := 3
	expect := 3
	runSample(t, arr, x, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 1}
	x := 4
	expect := 2
	runSample(t, arr, x, expect)
}
