package main

import "testing"

func runSample(t *testing.T, forest []string, expect int) {
	res := solve(forest)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample0(t *testing.T) {
	forest := []string{
		"aa",
		"aa",
	}
	expect := 2
	runSample(t, forest, expect)
}

func TestSample1(t *testing.T) {
	forest := []string{
		"aaab",
		"baaa",
		"abba",
	}
	expect := 3
	runSample(t, forest, expect)
}

func TestSample2(t *testing.T) {
	forest := []string{
		"aaaaaa",
		"aafaaa",
		"aaaafa",
		"aafaaa",
		"aaaaaa",
	}
	expect := 47
	runSample(t, forest, expect)
}
