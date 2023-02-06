package main

import "testing"

func runSample(t *testing.T, s string, k int, expect bool) {
	res := solve(s, k)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111"
	k := 2
	expect := true
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "10110111"
	k := 3
	expect := false
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "11111111"
	k := 4
	expect := true
	runSample(t, s, k, expect)
}
