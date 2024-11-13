package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0110"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0101"
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "00001000"
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0001000"
	expect := 3
	runSample(t, s, expect)
}
