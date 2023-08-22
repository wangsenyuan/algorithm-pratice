package main

import "testing"

func runSample(t *testing.T, x []string, expect string) {
	res := solve(x)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []string{
		"a",
		"ab",
		"b",
		"cd",
		"acd",
	}
	expect := "a"
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := []string{
		"z",
		"a",
		"a",
		"aa",
		"yakumo",
		"ran",
		"yakumoran",
	}
	expect := "z"
	runSample(t, x, expect)
}
