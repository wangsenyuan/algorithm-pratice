package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, r := 2, 47
	expect := 12
	runSample(t, l, r, expect)
}

func TestSample2(t *testing.T) {
	l, r := 47, 1024
	expect := 98
	runSample(t, l, r, expect)
}

func TestSample3(t *testing.T) {
	l, r := 1, 100
	// 1.... 9 = 9
	// 11    99 = 9
	expect := 18
	runSample(t, l, r, expect)
}
