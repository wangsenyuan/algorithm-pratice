package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else {
		var xor int
		for i := 1; i <= n; i++ {
			if res[i-1] == i {
				t.Fatalf("Sample result %v, not correct", res)
			}
			xor ^= abs(res[i-1] - i)
		}
		if xor != 0 {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample(t *testing.T) {
	for i := 4; i < 100; i++ {
		runSample(t, i, true)
	}
}
