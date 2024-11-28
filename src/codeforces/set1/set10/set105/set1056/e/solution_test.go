package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01"
	x := "aaaaaa"
	expect := 4
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "001"
	x := "kokokokotlin"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "001"
	x := "kokokokotlin"
	expect := 2
	runSample(t, s, x, expect)
}
