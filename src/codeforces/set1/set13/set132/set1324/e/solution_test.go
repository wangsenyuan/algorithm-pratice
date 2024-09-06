package main

import "testing"

func runSample(t *testing.T, h int, l int, r int, a []int, expect int) {
	res := solve(a, h, l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h, l, r := 24, 21, 23
	a := []int{16, 17, 14, 20, 20, 11, 22}
	expect := 3
	runSample(t, h, l, r, a, expect)
}
