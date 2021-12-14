package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(len(S), S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "NNN"
	expect := 2
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "PNP"
	expect := 0
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "NPNPNPNNP"
	expect := 2
	runSample(t, S, expect)
}
