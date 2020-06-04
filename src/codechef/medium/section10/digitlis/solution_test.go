package main

import "testing"

func runSample(t *testing.T, n int, L []int, expect int) {
	res := solve(n, L)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, L, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	L := []int{1, 2}
	expect := 36
	runSample(t, n, L, expect)
}

func TestSample2(t *testing.T) {
	n := 1
	L := []int{1}
	expect := 10
	runSample(t, n, L, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	L := []int{1, 2, 3}
	expect := 84
	runSample(t, n, L, expect)
}
