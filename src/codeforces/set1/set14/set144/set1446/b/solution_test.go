package main

import "testing"

func runSample(t *testing.T, A, B string, expect int) {
	res := solve(A, B)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "abba"
	B := "babab"
	expect := 5
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := "bbbbabab"
	B := "bbbabaaaaa"
	expect := 12
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := "uiibwws"
	B := "qhtkxcn"
	expect := 0
	runSample(t, A, B, expect)
}
