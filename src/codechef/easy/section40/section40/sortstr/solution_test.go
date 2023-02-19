package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cbaba"
	expect := "bbcaa"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "dbbced"
	expect := "cdbbde"
	runSample(t, s, expect)
}
