package main

import "testing"

func runSample(t *testing.T, n int, expect int64) {
	res := solve(n)
	if res != expect {
		t.Errorf("sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 18)
}
