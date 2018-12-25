package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int64) {
	res := solve(n, s)
	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, "xxx", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, "xxxyy", 11)
}
