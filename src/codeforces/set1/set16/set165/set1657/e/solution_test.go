package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 5)
}
func TestSample2(t *testing.T) {
	runSample(t, 4, 4, 571)
}
func TestSample5(t *testing.T) {
	runSample(t, 6, 9, 310640163)
}

func TestSample7(t *testing.T) {
	runSample(t, 42, 13, 136246935)
}
