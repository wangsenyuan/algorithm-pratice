package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 3, 3, 2
	expect := 6
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 2, 1, 10
	expect := 5
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 6, 3, 10
	expect := 375000012
	runSample(t, n, m, k, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 6, 4, 10
	expect := 500000026
	runSample(t, n, m, k, expect)
}

func TestSample5(t *testing.T) {
	n, m, k := 100, 1, 1
	expect := 958557139
	runSample(t, n, m, k, expect)
}
