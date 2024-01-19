package main

import "testing"

func runSample(t *testing.T, m int, k int, a1 int, ak int, expect int) {
	res := solve(m, k, a1, ak)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, k, a1, ak := 11, 3, 6, 1
	expect := 1
	runSample(t, m, k, a1, ak, expect)
}
