package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "DAAABDCA"
	expect := 11088
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "AB"
	expect := 10010
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ABCDEEDCBA"
	expect := 31000
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "DDDDAAADDABECD"
	expect := 15886
	runSample(t, s, expect)
}
