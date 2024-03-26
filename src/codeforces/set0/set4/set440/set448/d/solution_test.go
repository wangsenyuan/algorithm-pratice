package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 2, 2, 2
	expect := 2
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 2, 3, 4
	expect := 3
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 1, 10, 5
	expect := 5
	runSample(t, n, m, k, expect)
}
