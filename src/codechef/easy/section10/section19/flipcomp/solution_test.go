package main

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := solve(S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "00110101100"
	expect := 4
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "100011"
	expect := 2
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "1110011"
	expect := 2
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "000101110"
	expect := 3
	runSample(t, S, expect)
}
