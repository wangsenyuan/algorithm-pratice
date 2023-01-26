package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	expect := 0
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 3
	expect := 4
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 4, 1
	expect := 2
	runSample(t, n, m, expect)
}
