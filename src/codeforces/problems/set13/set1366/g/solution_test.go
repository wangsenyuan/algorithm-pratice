package main

import "testing"

func runSample(t *testing.T, s, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "a.ba.b."
	b := "abb"
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := ".bbac..a.c.cd"
	b := "bacd"
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "c..code..c...o.d.de"
	b := "code"
	expect := 3
	runSample(t, a, b, expect)
}
