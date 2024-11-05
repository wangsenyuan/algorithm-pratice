package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	m := len(res)

	for i := 0; i < m; i++ {
		var sum int
		for j := i; j < m; j++ {
			sum ^= res[j]
		}
		if sum == 0 || sum == x {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 0)
}
