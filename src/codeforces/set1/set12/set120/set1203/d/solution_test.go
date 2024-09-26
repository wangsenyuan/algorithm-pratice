package main

import "testing"

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "bbaba"
	x := "bb"
	expect := 3
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "bbaba"
	x := "ab"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "abcde"
	x := "abcde"
	expect := 0
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "asdfasdf"
	x := "fasd"
	expect := 3
	runSample(t, s, x, expect)
}
