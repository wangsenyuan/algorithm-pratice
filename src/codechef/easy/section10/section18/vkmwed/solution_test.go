package main

import "testing"

func runSample(t *testing.T, m, n int, expect int) {
	res := solve(m, n)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 7)
}

func TestSample2(t *testing.T) {
	runSample(t, 12, 6, 255)
}

func TestSample3(t *testing.T) {
	runSample(t, 6969, 8563, 988402422)
}
