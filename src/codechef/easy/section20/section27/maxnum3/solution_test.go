package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "123"
	expect := "12"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "100222"
	expect := "00222"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "10003"
	expect := "-1"
	runSample(t, s, expect)
}
