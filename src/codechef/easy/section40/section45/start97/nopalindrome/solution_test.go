package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(k, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 1
	expect := 1
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	expect := 1
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 2
	// 10021
	expect := 4
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 7, 4
	// 1000020
	expect := 3
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 8, 3
	// 10002010
	expect := 4
	runSample(t, n, k, expect)
}
