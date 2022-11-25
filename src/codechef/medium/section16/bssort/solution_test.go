package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := int(solve(s))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0011"
	expect := 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111000"
	expect := 17
	runSample(t, s, expect)
}
