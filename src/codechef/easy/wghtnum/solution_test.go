package main

import "testing"

func runSample(t *testing.T, n int64, w int, expect int64) {
	res := solve(n, w)
	if res != expect {
		t.Errorf("sample %d %d, expect %d, but got %d", n, w, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 6)
}
