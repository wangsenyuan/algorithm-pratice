package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "abacaba"
	b := "bacabaa"
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "zcabd"
	b := "dbacz"
	expect := 0
	runSample(t, a, b, expect)
}
