package main

import "testing"

func runSample(t *testing.T, P, Q int, S string, expect bool) {
	res := solve(len(S), P, Q, S)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "111"
	P, Q := 0, -1
	expect := true
	runSample(t, P, Q, S, expect)
}


func TestSample2(t *testing.T) {
	S := "1101"
	P, Q := 2, 0
	expect := true
	runSample(t, P, Q, S, expect)
}

func TestSample3(t *testing.T) {
	S := "00"
	P, Q := 0, 1
	expect := false
	runSample(t, P, Q, S, expect)
}
