package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 12
	expect := 7
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 5
	expect := 3
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 4, 4
	expect := 3
	runSample(t, n, m, expect)
}
