package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "7474"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "47"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "47744"
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "7744"
	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "474747447444474"
	expect := 11
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "44"
	expect := 2
	runSample(t, s, expect)
}
