package main

import "testing"

func runSample(t *testing.T, S, P string, expect bool) {
	res := solve(S, P)

	if res != expect {
		t.Errorf("Sample %s %s, expect %t, but got %t", S, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "00"
	P := "00"
	expect := true
	runSample(t, S, P, expect)
}

func TestSample2(t *testing.T) {
	S := "101"
	P := "010"
	expect := false
	runSample(t, S, P, expect)
}

func TestSample3(t *testing.T) {
	S := "0110"
	P := "0011"
	expect := true
	runSample(t, S, P, expect)
}
