package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1010"
	k := 0
	expect := 21
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "0010100"
	k := 1
	expect := 22
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "00110"
	k := 2
	expect := 12
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "0100"
	k := 1
	expect := 10
	runSample(t, s, k, expect)
}
