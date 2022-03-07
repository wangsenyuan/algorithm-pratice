package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if expect != (len(res) > 0) {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else if expect {
		var xor int
		for i := 1; i < n; i++ {
			xor ^= abs(res[i] - res[i-1])
		}
		if xor != 0 {
			t.Errorf("Sample result %v, not correct, xor got %d", res, xor)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	runSample(t, 6, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, true)
}
