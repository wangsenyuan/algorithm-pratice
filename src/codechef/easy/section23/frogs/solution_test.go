package main

import "testing"

func runSample(t *testing.T, n int, W, L []int, expect int) {
	res := solve(n, W, L)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	W := []int{3, 1, 2}
	L := []int{1, 4, 5}
	expect := 3
	runSample(t, n, W, L, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	W := []int{3, 2, 1}
	L := []int{1, 1, 1}
	expect := 6
	runSample(t, n, W, L, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	W := []int{2, 1, 4, 3}
	L := []int{4, 1, 2, 4}
	expect := 5
	runSample(t, n, W, L, expect)
}
