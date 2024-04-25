package main

import "testing"

func runSample(t *testing.T, k int, a string, b string, expect int) {
	res := solve(k, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "aa"
	b := "bb"
	k := 4
	expect := 6
	runSample(t, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "abbb"
	b := "baaa"
	k := 5
	expect := 8
	runSample(t, k, a, b, expect)
}
