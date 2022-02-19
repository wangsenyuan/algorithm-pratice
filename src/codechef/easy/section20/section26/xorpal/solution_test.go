package main

import "testing"

func runSample(t *testing.T, S string, expect bool) {
	res := solve(len(S), S)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "00"
	expect := true
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "0011"
	expect := true
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "001"
	expect := true
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "0001"
	expect := false
	runSample(t, S, expect)
}
