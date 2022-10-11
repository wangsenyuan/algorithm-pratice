package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ab"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aa"
	expect := 1
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := "abb"
	expect := 2
	runSample(t, s, expect)
}
