package main

import "testing"

func runSample(t *testing.T, a int, b int, s string, expect int) {
	res := solve(a, b, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 2, 5
	s := "00110010"
	expect := 94
	runSample(t, a, b, s, expect)
}

func TestSample2(t *testing.T) {
	a, b := 1, 1
	s := "00110010"
	expect := 25
	runSample(t, a, b, s, expect)
}

func TestSample3(t *testing.T) {
	a, b := 100000000, 100000000
	s := "010101010"
	expect := 2900000000
	runSample(t, a, b, s, expect)
}

func TestSample4(t *testing.T) {
	a, b := 5, 1
	s := "00"
	expect := 13
	runSample(t, a, b, s, expect)
}
