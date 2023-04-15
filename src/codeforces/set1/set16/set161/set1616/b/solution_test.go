package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sampl expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "baa"
	expect := "baaaab"
	runSample(t, s, expect)
}
