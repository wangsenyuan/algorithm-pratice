package main

import "testing"

func runSample(t *testing.T, n, x, y int, expect int) {
	res := int(solve(n, x, y))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSaample1(t *testing.T) {
	runSample(t, 9, 40, 8, 24)
}

func TestSaample2(t *testing.T) {
	runSample(t, 15, 20, 20, 20)
}

func TestSaample3(t *testing.T) {
	runSample(t, 105, 80, 10, 100)
}
