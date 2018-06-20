package main

import "testing"

func runSample(t *testing.T, x, y string, a, b, k int, expect int) {
	res := solve([]byte(x), []byte(y), a, b, k)

	if res != expect {
		t.Errorf("Sample %s %s %d %d %d, expect %d, but got %d", x, y, a, b, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aaa", "bbbb", 0, 0, 100, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "abab", "acac", 1, 1, 100, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "baaaaa", "aaaaab", 1, 100, 100, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, "aaaaaa", "bbbbbb", 100, 100, 0, -1)
}
