package main

import "testing"

func runSample(t *testing.T, n, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 2115)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 350, 895852507)
}
