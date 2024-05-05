package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, expect int) {
	res := solve(n, a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 8, 1, 2
	expect := 1
	runSample(t, n, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 5, 3, 4
	expect := 6
	runSample(t, n, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 6, 4, 2
	expect := 4
	runSample(t, n, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, a, b := 20, 5, 6
	expect := 2
	runSample(t, n, a, b, expect)
}
