package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(S)

	if res != expect {
		t.Errorf("Sample expect %d, bug got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "000111"
	expect := 1
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "11100"
	expect := 2
	runSample(t, S, expect)
}
