package main

import "testing"

func runSample(t *testing.T, s string, x string, expect bool) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "sofia"
	x := "afios"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "cba"
	x := "bc"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "sofia"
	x := "e"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "anavolimilovana"
	x := "aamanan"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample5(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	x := "nope"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample6(t *testing.T) {
	s := "zyxwvutsrqponmlkjihgfedcba"
	x := "nope"
	expect := true
	runSample(t, s, x, expect)
}

func TestSample7(t *testing.T) {
	s := "apricot"
	x := "cat"
	expect := false
	runSample(t, s, x, expect)
}

func TestSample8(t *testing.T) {
	s := "cba"
	x := "acb"
	expect := true
	runSample(t, s, x, expect)
}
