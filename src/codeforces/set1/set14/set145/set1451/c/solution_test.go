package main

import "testing"

func runSample(t *testing.T, a string, b string, k int, expect bool) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "abc"
	b := "bcd"
	k := 3
	expect := false
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := "abba"
	b := "azza"
	k := 2
	expect := true
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a := "zz"
	b := "aa"
	k := 1
	expect := false
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	a := "aaabba"
	b := "ddddcc"
	k := 2
	expect := true
	runSample(t, a, b, k, expect)
}
