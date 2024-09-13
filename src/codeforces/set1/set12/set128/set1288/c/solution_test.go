package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 1, 55)
}

func TestSample3(t *testing.T) {
	runSample(t, 723, 9, 157557417)
}

func TestSample4(t *testing.T) {
	// 1,1 vs 1 1
	// 1,1 vs 2,1
	// 1,1 vs 3,1
	// 1,2, vs 2 2
	// 1,2 vs 3,2
	// 1,2 vs 3,3
	runSample(t, 3, 2, 15)
}
