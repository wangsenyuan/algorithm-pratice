package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 22, 4
	expect := 4
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 9
	expect := 0
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 55, 11
	expect := 1
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	n, k := 42069, 6
	expect := 1052
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 69420, 4
	expect := 11571
	runSample(t, n, k, expect)
}
