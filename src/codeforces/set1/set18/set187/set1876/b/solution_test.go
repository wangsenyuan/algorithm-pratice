package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{19, 14, 19, 9}
	expect := 265
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{90000, 9000, 99000, 900, 90900, 9900, 99900, 90, 90090, 9090, 99090, 990, 90990, 9990, 99990}
	expect := 266012571
	runSample(t, a, expect)
}
