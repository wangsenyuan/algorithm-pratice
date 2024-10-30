package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 3, 5, 6}
	b := []int{11, 12, 14, 20}
	k := 3
	expect := 3
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 4}
	b := []int{11, 3, 16}
	k := 1
	expect := 4
	runSample(t, a, b, k, expect)
}
