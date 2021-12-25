package main

import "testing"

func runSample(t *testing.T, X int, expect int) {
	res := solve(X)

	if res != expect {
		t.Errorf("sample %d, expect %d, but got %d", X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 9, 5)
}
