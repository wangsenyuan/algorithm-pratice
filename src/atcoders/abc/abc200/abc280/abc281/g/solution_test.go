package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 1000000000
	expect := 8
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 1000000000
	expect := 1
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 500, 987654321
	expect := 610860515
	runSample(t, n, m, expect)
}
