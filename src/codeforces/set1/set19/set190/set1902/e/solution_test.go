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
		"aba",
		"ab",
		"ba",
	}
	expect := 20
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"abab",
		"babx",
		"xab",
		"xba",
		"bab",
	}
	expect := 126
	runSample(t, words, expect)
}
