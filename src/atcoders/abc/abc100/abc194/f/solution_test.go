package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(k, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10"
	k := 1
	expect := 15
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "FF"
	k := 2
	expect := 225
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "100"
	k := 2
	expect := 226
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "110"
	k := 2
	expect := 225 + 1 + 1 + 1
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "1A8FD02"
	k := 4
	expect := 3784674
	runSample(t, s, k, expect)
}
