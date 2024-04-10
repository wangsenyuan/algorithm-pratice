package main

import "testing"

func runSample(t *testing.T, k int, a []int, expect int) {
	res := solve(k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{50, 110, 130, 40, 120}
	expect := 20
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	a := []int{2, 3, 4, 1}
	expect := 0
	runSample(t, k, a, expect)
}
