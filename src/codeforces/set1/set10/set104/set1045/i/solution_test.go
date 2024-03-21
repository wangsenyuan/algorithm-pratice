package main

import (
	"testing"
)

func runSample(t *testing.T, words []string, expect int) {
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"aa",
		"bb",
		"cd",
	}
	expect := 1
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"aab",
		"abcac",
		"dffe",
		"ed",
		"aa",
		"aade",
	}
	expect := 6
	runSample(t, words, expect)
}
