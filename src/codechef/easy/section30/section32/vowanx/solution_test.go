package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcdefghij"
	expect := "hgfeabcdij"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "bcadage"
	expect := "gacbade"
	runSample(t, s, expect)
}