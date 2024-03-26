package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "1000"
	b := "111"
	expect := -1
	runSample(t, a, b, expect)
}
func TestSample2(t *testing.T) {
	a := "00100"
	b := "11"
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "110"
	b := "101"
	expect := 1
	runSample(t, a, b, expect)
}
