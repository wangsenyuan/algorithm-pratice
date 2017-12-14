package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)
	if res != expect {
		t.Errorf("fast square of %d ** %d, should give %d, but got %d", x, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4, 139)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 6, 40079781)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 5, 32745632)
}
