package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0000"
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "01010101"
	expect := 130
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "01"
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "101"
	// (0, 1) 1
	// (0, 2) 2
	// (1, 2) 1
	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "1100111001"
	expect := 147
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "11000000111"
	expect := 70
	runSample(t, s, expect)
}
