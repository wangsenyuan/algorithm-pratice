package main

import "testing"

func runSample(t *testing.T, X int64, c int, expect int64) {
	res := solve(X, c)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 5, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 1, 1)
}


func TestSample3(t *testing.T) {
	runSample(t, 20, 4, 20)
}