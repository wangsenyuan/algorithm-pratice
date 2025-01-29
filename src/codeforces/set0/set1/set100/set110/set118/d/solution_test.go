package main

import "testing"

func runSample(t *testing.T, n1 int, n2 int, k1 int, k2 int, expect int) {
	res := solve(n1, n2, k1, k2)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n1, n2, k1, k2 := 2, 1, 1, 10
	expect := 1
	runSample(t, n1, n2, k1, k2, expect)
}

func TestSample2(t *testing.T) {
	n1, n2, k1, k2 := 2, 3, 1, 2
	expect := 5
	runSample(t, n1, n2, k1, k2, expect)
}

func TestSample3(t *testing.T) {
	n1, n2, k1, k2 := 2, 4, 1, 1
	expect := 0
	runSample(t, n1, n2, k1, k2, expect)
}
