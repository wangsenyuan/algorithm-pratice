package main

import "testing"

func runSample(t *testing.T, k int, words []string, expect string) {
	res := solve(k, words)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	words := []string{
		"a", "b",
	}
	expect := "First"
	runSample(t, k, words, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	words := []string{
		"a", "b", "c",
	}
	expect := "First"
	runSample(t, k, words, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	words := []string{
		"ab",
	}
	expect := "Second"
	runSample(t, k, words, expect)
}

func TestSample4(t *testing.T) {
	k := 6
	words := []string{
		"abas",
		"dsfdf",
		"abacaba",
		"dartsidius",
		"kolobok",
	}
	expect := "Second"
	runSample(t, k, words, expect)
}
