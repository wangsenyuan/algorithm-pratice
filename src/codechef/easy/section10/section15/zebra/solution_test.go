package main

import "testing"

func runSample(t *testing.T, k int, S string, expect int) {
	res := solve(len(S), k, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	S := "100101"
	expect := 6
	runSample(t, k, S, expect)
}


func TestSample2(t *testing.T) {
	k := 1
	S := "10111"
	expect := 2
	runSample(t, k, S, expect)
}


func TestSample3(t *testing.T) {
	k := 1
	S := "000000"
	expect := -1
	runSample(t, k, S, expect)
}
