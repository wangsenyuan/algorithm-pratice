package main

import "testing"

func runSample(t *testing.T, a string, b string, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := "abc"
	b := "abd"
	expect := 6
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	k := 0
	a := "abc"
	b := "abd"
	expect := 3
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	a := "xbb"
	b := "xcd"
	expect := 6
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	a := "lkwhbahuqa"
	b := "qoiujoncjb"
	expect := 11
	runSample(t, a, b, k, expect)
}
