package main

import "testing"

func runSample(t *testing.T, x, y string, ln int, cnt int64) {
	a, b := solve([]byte(x), []byte(y))

	if a != ln || b != cnt {
		t.Errorf("Sample %s %s, expect %d %d, but got %d %d", x, y, ln, cnt, a, b)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "acbd", "acbd", 4, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "vnvn", "vn", 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "abcd", "abdc", 3, 2)
}
