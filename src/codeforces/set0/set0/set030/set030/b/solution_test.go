package main

import "testing"

func runSample(t *testing.T, a string, b string, expect bool) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "01.01.98"
	b := "01.01.80"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "20.10.20"
	b := "10.02.30"
	expect := false
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "28.02.74"
	b := "28.02.64"
	expect := false
	runSample(t, a, b, expect)
}
