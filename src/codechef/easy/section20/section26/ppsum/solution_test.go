package main

import "testing"

func runSample(t *testing.T, d int, n int, expect int) {
	res := solve(d, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d, n := 1, 4
	expect := 10
	runSample(t, d, n, expect)
}

func TestSample2(t *testing.T) {
	d, n := 2, 3
	expect := 21
	runSample(t, d, n, expect)
}

