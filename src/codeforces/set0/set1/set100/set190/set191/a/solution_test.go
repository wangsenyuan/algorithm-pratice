package main

import "testing"

func runSample(t *testing.T, s []string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"abc",
		"ca",
		"cba",
	}
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"vvp",
		"vvp",
		"dam",
		"vvp",
	}
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []string{
		"ab",
		"c",
		"def",
	}
	expect := 1
	runSample(t, s, expect)
}
