package main

import "testing"

func runSample(t *testing.T, L int, R int, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 20, 35)
}
