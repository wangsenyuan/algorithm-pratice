package main

import "testing"

func runSample(t *testing.T, words []string, expect string) {
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"mail",
		"ai",
		"lru",
		"cf",
	}
	expect := "cfmailru"
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"kek",
		"preceq",
		"cheburek",
	}
	expect := "NO"
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"z",
	}
	expect := "z"
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{
		"ac",
		"bc",
	}
	expect := "NO"
	runSample(t, words, expect)
}
