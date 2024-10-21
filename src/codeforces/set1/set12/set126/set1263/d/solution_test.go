package main

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"a",
		"b",
		"ab",
		"d",
	}
	expect := 2
	runSample(t, words, expect)
}
