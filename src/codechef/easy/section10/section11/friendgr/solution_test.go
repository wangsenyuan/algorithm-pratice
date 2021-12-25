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
	S := "010010"
	expect := 1
	runSample(t, k, S, expect)
}


func TestSample2(t *testing.T) {
	k := 0
	S := "1001"
	expect := 2
	runSample(t, k, S, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	S := "1101001"
	expect := 1
	runSample(t, k, S, expect)
}
