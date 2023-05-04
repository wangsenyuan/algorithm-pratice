package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	var expect int64 = 7
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 4
	var expect int64 = 5
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 4, 2
	var expect int64 = 6
	runSample(t, n, m, expect)
}
