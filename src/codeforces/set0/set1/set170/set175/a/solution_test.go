package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1234"
	expect := 37
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "9000"
	expect := 90
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0009"
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "100000010000011000000"
	expect := -1
	runSample(t, s, expect)
}
