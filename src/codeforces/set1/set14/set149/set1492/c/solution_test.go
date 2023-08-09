package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abbbc"
	x := "abc"
	expect := 3
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "aaaaa"
	x := "aa"
	expect := 4
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "abcdf"
	x := "abcdf"
	expect := 1
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "ab"
	x := "ab"
	expect := 1
	runSample(t, s, x, expect)
}
