package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}
	expect := 13
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{3, 2, 1, 1, 4, 5, 1}
	expect := 6
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 2}
	expect := 5
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{3, 3, 9, 10, 1, 4, 10, 2, 9, 4, 2, 8, 8, 3, 9, 7, 1, 3, 8, 8, 8, 8, 2, 7, 9, 8, 2}
	expect := 13
	runSample(t, arr, expect)
}
