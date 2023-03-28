package main

import "testing"

func runSample(t *testing.T, n int) {
	var h int
	for (1 << (h + 1)) < n {
		h++
	}
	expect := 1 << h

	res := solve(n)

	x := res[0] ^ res[1]
	for i := 1; i+1 < n; i++ {
		x = max(x, res[i]^res[i+1])
	}

	if x != expect {
		t.Fatalf("Sample %d expec %d, but got %v with max xor %d", n, expect, res, x)
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestSample1(t *testing.T) {
	for i := 2; i <= 1000; i++ {
		runSample(t, i)
	}
}

func TestSample2(t *testing.T) {
	runSample(t, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 6)
}
