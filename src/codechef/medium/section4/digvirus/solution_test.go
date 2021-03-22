package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(len(S), []byte(S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "555755555"
	expect := 3
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "311110000000000"
	expect := 6
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "07788000744"
	expect := 4
	runSample(t, S, expect)
}
