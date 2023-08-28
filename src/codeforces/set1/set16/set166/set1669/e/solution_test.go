package main

import "testing"

func runSample(t *testing.T, s []string, expect int) {
	res := int(solve(s))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"ab",
		"cb",
		"db",
		"aa",
		"cc",
		"ef",
	}
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"aa",
		"bb",
		"cc",
		"ac",
		"ca",
		"bb",
		"aa",
	}
	expect := 6
	runSample(t, s, expect)
}
