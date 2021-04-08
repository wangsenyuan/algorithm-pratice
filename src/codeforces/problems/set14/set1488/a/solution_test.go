package main

import "testing"

func runSample(t *testing.T, x, y int, expect int) {
	res := solve(x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 7, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 42, 5)
}

func TestSample4(t *testing.T) {
	runSample(t, 25, 1337, 20)
}
