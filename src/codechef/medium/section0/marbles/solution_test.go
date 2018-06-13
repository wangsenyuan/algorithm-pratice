package main

import "testing"

func runSample(t *testing.T, n, k int, expect int64) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 10, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 30, 7, 475020)
}
