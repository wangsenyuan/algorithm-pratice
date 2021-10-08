package main

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 26)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 10000000000000000, 366580019)
}
