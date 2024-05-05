package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 3126690179932000, 2474382898739836, -1
	expect := 917305624
	runSample(t, n, m, k, expect)
}
