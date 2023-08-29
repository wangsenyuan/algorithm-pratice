package main

import "testing"

func runSample(t *testing.T, candies []int, expect int) {
	res := solve(candies)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	candies := []int{2133, 1012, 371, 6963, 1941, 1458, 47}
	expect := 0
	runSample(t, candies, expect)
}
