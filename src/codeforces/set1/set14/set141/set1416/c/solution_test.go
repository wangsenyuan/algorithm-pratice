package main

import "testing"

func runSample(t *testing.T, a []int, expect []int) {
	cnt, x := solve(a)

	if expect[0] != cnt || expect[1] != x {
		t.Fatalf("Sample expect %v, but got %d %d", expect, x, cnt)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1, 3, 2}
	expect := []int{1, 0}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 7, 9, 10, 7, 5, 5, 3, 5}
	expect := []int{4, 14}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{8, 10, 3}
	expect := []int{0, 8}
	runSample(t, a, expect)
}
