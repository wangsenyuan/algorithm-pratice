package main

import "testing"

func runSample(t *testing.T, a string, b string, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "aab"
	b := "bcc"
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "cabc"
	b := "abcb"
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "abc"
	b := "tsr"
	expect := 3
	runSample(t, a, b, expect)
}
func TestSample4(t *testing.T) {
	a := "aabd"
	b := "cccd"
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := "abcbd"
	b := "bcdda"
	expect := 4
	runSample(t, a, b, expect)
}
