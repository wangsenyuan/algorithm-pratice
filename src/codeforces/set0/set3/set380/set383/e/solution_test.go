package main

import "testing"

func runSample(t *testing.T, words []string, expect int64) {
	res := solve(words)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"abc",
		"aaa",
		"ada",
		"bcd",
		"def",
	}
	var expect int64
	runSample(t, words, expect)
}
