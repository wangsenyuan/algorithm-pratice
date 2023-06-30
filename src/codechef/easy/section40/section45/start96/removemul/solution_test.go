package main

import "testing"

func runSample(t *testing.T, n int, Q []int, expect int) {
	res := int(solve(n, Q))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {

}
