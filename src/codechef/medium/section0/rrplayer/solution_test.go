package main

import "testing"

func runSample(t *testing.T, n int, expect float64) {
	res := solve(n)
	if res != expect {
		t.Errorf("sample %d expect %f, but got %f", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1.0)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3.0)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 5.5)
}
