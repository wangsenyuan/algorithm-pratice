package main

import "testing"

func runSample(t *testing.T, n int, xs []int, expect int) {
	res := solve(n, xs)
	if res != expect {
		t.Errorf("sample: %d %v, should give %d, but got %d", n, xs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	xs := []int{2, 4, 1}
	expect := 6
	runSample(t, n, xs, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	xs := []int{2, 4, 1, 3}
	expect := 13
	runSample(t, n, xs, expect)
}
