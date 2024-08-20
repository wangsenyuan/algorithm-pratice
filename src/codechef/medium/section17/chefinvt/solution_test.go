package main

import "testing"

func runSample(t *testing.T, n int, p int, k int, expect int) {
	res := solve(n, p, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p, k := 10, 5, 5
	expect := 2
	runSample(t, n, p, k, expect)
}

func TestSample2(t *testing.T) {
	n, p, k := 10, 6, 5
	expect := 4
	runSample(t, n, p, k, expect)
}

func TestSample3(t *testing.T) {
	n, p, k := 10, 4, 5
	expect := 9
	runSample(t, n, p, k, expect)
}

func TestSample4(t *testing.T) {
	n, p, k := 10, 8, 5
	expect := 8
	runSample(t, n, p, k, expect)
}
