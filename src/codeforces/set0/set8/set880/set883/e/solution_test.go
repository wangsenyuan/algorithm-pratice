package main

import (
	"testing"
)

func runSample(t *testing.T, s string, words []string, expect int) {
	res := solve(s, words)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a**d"
	words := []string{
		"abcd",
		"acbd",
	}
	expect := 2
	runSample(t, s, words, expect)
}

func TestSample2(t *testing.T) {
	s := "lo*er"
	words := []string{
		"lover",
		"loser",
	}
	expect := 0
	runSample(t, s, words, expect)
}
func TestSample3(t *testing.T) {
	s := "a*a"
	words := []string{
		"aaa",
		"aba",
	}
	expect := 1
	runSample(t, s, words, expect)
}
