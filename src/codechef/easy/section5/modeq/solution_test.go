package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 6
	var expect int64 = 3
	runSample(t, n, m, expect)
}
