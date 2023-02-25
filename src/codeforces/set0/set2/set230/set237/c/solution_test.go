package main

import "testing"

func runSample(t *testing.T, a, b int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Errorf("Smaple expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b, k := 2, 4, 2
	expect := 3
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a, b, k := 6, 13, 1
	expect := 4
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a, b, k := 1, 4, 3
	expect := -1
	runSample(t, a, b, k, expect)
}
