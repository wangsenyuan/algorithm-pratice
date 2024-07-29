package main

import "testing"

func runSample(t *testing.T, a, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "100"
	b := "001"
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "0101"
	b := "0011"
	expect := 1
	runSample(t, a, b, expect)
}
