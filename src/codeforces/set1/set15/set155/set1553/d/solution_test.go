package main

import "testing"

func runSample(t *testing.T, s string, x string, expect bool) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %t,b ut got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababa"
	x := "ba"
	expect := true
	runSample(t, s, x, expect)
}
