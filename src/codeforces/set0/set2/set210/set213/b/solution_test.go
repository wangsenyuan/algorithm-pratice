package main

import "testing"

func runSample(t *testing.T, n int, a []int, expect int) {
	res := solve(n, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	// 10
	res := bruteForce(n, a)

	t.Logf("brute force %v", res)

	expect := len(res)
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	a := []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	expect := 1
	runSample(t, n, a, expect)
}
