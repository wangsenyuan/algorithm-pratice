package main

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample %v expect %d, but got %d", words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"snuke",
		"kensho",
		"uk",
	}
	expect := 9
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"abc",
		"abc",
		"arc",
	}
	expect := 6
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"cmcmrcc",
		"rmrrrmr",
		"mrccm",
		"mmcr",
		"rmmrmrcc",
		"ccmcrcmcm",
	}
	expect := 27
	runSample(t, words, expect)
}
