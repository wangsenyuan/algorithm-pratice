package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3, 294)
}
func TestSample3(t *testing.T) {
	runSample(t, 2020, 2021, 50657649)
}
