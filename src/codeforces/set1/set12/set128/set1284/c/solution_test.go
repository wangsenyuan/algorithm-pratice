package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 993244853, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 993244853, 32)
}

func TestSample3(t *testing.T) {
	runSample(t, 2019, 993244853, 923958830)
}
