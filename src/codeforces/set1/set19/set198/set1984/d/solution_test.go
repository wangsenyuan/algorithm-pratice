package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaaa"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "baba"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "cabacb"
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "aaabaaa"
	expect := 16
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "bitset"
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "ab"
	expect := 2
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "abbaaaabbb"
	expect := 3
	runSample(t, s, expect)
}
func TestSample8(t *testing.T) {
	s := "yearnineteeneightyfour"
	expect := 1
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "baaab"
	expect := 2
	runSample(t, s, expect)
}
