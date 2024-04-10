package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aabaac"
	expect := 2
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := "0rTrT022"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aA"
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "TddTddddTddddddTdddTdddddddddddddddddddd"
	expect := 8
	runSample(t, s, expect)
}
