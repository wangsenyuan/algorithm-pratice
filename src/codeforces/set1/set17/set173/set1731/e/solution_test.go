package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int64) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 1
	var expect int64 = 2
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 6, 10
	var expect int64 = -1
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 9, 4
	var expect int64 = 7
	runSample(t, n, m, expect)
}

func TestSample4(t *testing.T) {
	n, m := 10, 11
	var expect int64 = 21
	runSample(t, n, m, expect)
}
