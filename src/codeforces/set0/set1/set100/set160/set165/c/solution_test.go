package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	s := "1010"
	expect := 6
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	s := "01010"
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	k := 100
	s := "01010"
	expect := 0
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	k := 0
	s := "01010"
	expect := 3
	runSample(t, s, k, expect)
}
