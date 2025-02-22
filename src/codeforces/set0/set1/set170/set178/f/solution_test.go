package main

import "testing"

func runSample(t *testing.T, k int, words []string, expect int) {
	res := solve(k, words)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	words := []string{
		"aba",
		"bzd",
		"abq",
	}
	expect := 2
	runSample(t, k, words, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	words := []string{
		"eee",
		"rrr",
		"ttt",
		"qqq",
	}
	expect := 0
	runSample(t, k, words, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	words := []string{
		"aaa",
		"abba",
		"abbc",
		"abbd",
	}
	expect := 9
	runSample(t, k, words, expect)
}

func TestSample4(t *testing.T) {
	k := 2
	words := []string{
		"goodbyebye",
		"goodbye",
	}
	expect := 7
	runSample(t, k, words, expect)
}

func TestSample5(t *testing.T) {
	k := 3
	words := []string{
		"a",
		"ab",
		"abc",
		"abcd",
		"abcde",
	}
	expect := 10
	runSample(t, k, words, expect)
}

func TestSample6(t *testing.T) {
	k := 5
	words := []string{
		"a",
		"handful",
		"of",
		"strings",
		"my",
		"hands",
		"are",
		"full",
		"of",
		"strings",
	}
	expect := 11
	runSample(t, k, words, expect)
}

func TestSample7(t *testing.T) {
	k := 10
	words := []string{
		"commonprefixhello",
		"commonprefixhow",
		"commonprefixdo",
		"commonprefixyou",
		"commonprefixdo",
		"commonprefixaaabd",
		"commonprefixaaab",
		"commonprefixabddef",
		"commonprefixbaaaae",
		"commonprefixbaae",
		"commonprefixbaaade",
		"commonprefixaedbed",
		"commonprefixaeeee",
		"commonprefixbbbbb",
		"commonprefixbebe",
		"commonprefixbebeb",
		"commonprefixbye",
		"commonprefixsee",
		"commonprefixyou",
		"commonprefixlater",
	}
	expect := 577
	runSample(t, k, words, expect)
}
