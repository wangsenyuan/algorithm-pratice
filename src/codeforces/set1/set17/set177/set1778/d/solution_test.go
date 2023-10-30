package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "0"
	b := "1"
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "1000"
	b := "1110"
	expect := 665496254
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "01001"
	b := "10111"
	expect := 665496277
	runSample(t, a, b, expect)
}
