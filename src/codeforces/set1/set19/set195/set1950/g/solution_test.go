package main

import "testing"

func runSample(t *testing.T, songes []string, expect int) {
	res := solve(songes)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	songes := []string{
		"electronic themotans",
		"electronic carlasdreams",
		"pop themotans",
		"pop irinarimes",
	}
	expect := 0
	runSample(t, songes, expect)
}
