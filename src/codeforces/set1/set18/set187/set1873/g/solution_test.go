package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ABBA"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "ABA"
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "BAABA"
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "ABB"
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "AAAAAAB"
	expect := 6
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "BABA"
	expect := 2
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "B"
	expect := 0
	runSample(t, s, expect)
}
