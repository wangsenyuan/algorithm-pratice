package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "1010"
	b := "1101"
	expect := 12
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "1001"
	b := "10101"
	expect := 11
	runSample(t, a, b, expect)
}
