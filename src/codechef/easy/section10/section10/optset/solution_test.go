package main

import "testing"

func runSample(t *testing.T, k int, n int, expectXor int) {
	res := solve(k, n)
	var xor int
	for _, cur := range res {
		xor ^= cur
	}

	if xor != expectXor {
		t.Errorf("Sample expect %d, but got %d", expectXor, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 9, 15)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, 8, 15)
}
