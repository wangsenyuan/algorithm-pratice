package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int64) {
	nums, ops := parse(n, []byte(s))
	res := solve(n, nums, ops)

	if res != expect {
		t.Errorf("Sample %d %s, expect %d, but got %d", n, s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, "8 - 1 - 25", 32)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, "38 + 20", 58)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, "40 - 8 - 1 + 25", 58)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, "10 - 40 - 8 - 1 + 25", 4)
}
