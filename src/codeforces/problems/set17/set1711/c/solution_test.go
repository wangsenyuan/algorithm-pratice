package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	expect := 12
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1110"
	expect := 780
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11011111101010010"
	expect := 141427753
	runSample(t, s, expect)
}
