package main

import "testing"

func runSample(t *testing.T, n int, m int, a int, b int, expect int) {
	res := solve(n, m, a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, a, b := 20, 5, 2, 20
	expect := 2
	runSample(t, n, m, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, m, a, b := 21, 5, 1, 13
	expect := 2
	runSample(t, n, m, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, m, a, b := 19, 5, 7, 19
	expect := 2
	runSample(t, n, m, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, m, a, b := 21, 3, 6, 11
	expect := 2
	runSample(t, n, m, a, b, expect)
}
