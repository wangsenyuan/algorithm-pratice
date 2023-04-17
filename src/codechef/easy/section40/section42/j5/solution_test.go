package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "difference"
	expect := "difference"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "similarq"
	expect := "imilaq"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "caaat"
	expect := "aat"
	runSample(t, s, expect)
}
