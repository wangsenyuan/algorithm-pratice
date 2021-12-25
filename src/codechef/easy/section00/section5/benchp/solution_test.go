package main

import "testing"

func runSample(t *testing.T, n, w, wr int, W []int, expect bool) {
	res := solve(n, w, wr, W)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, w, wr := 7, 100, 50
	W := []int{100, 10, 10, 10, 10, 10, 90}
	expect := false
	runSample(t, n, w, wr, W, expect)
}

func TestSample2(t *testing.T) {
	n, w, wr := 6, 100, 40
	W := []int{10, 10, 10, 10, 10, 10}
	expect := true
	runSample(t, n, w, wr, W, expect)
}
