package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(len(s), s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abdeba"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "codechef"
	expect := 7
	runSample(t, s, expect)
}
