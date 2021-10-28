package main

import "testing"

func runSample(t *testing.T, n, v int, a, b int64) {
	c, d := solve(n, v)

	if c != a || d != b {
		t.Errorf("Sample expect %d, %d, but got %d %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	n, v := 3, 1
	var a, b int64 = 3, 3
	runSample(t, n, v, a, b)
}

func TestSample2(t *testing.T) {
	n, v := 4, 2
	var a, b int64 = 6, 4
	runSample(t, n, v, a, b)
}


func TestSample3(t *testing.T) {
	n, v := 4, 3
	var a, b int64 = 6, 3
	runSample(t, n, v, a, b)
}
