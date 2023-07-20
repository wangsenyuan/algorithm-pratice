package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0010"
	expect := "0110"
	runSample(t, s, expect)
}
