package main

import "testing"

func runSample(t *testing.T, a, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "100001"
	b := "110111"
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "000"
	b := "111"
	expect := 2
	runSample(t, a, b, expect)
}
