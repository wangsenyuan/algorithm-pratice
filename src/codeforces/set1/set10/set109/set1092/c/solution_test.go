package main

import "testing"

func TestSample1(t *testing.T) {
	n := 5
	words := []string{
		"ba",
		"a",
		"abab",
		"a",
		"aba",
		"baba",
		"ab",
		"aba",
	}
	res := solve(n, words)
	if len(res) != len(words) {
		t.Fatalf("Sample result %s, not correct", res)
	}
}
