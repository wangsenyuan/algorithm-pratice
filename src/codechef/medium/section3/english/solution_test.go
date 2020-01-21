package main

import "testing"

func runSample(t *testing.T, n int, words []string, expect int64) {
	res := solve(n, words)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	words := []string{"abcdefghijkl",
		"chef",
		"world",
		"code",
		"abcxyzhijkl",
		"word",
	}
	expect := int64(10)

	runSample(t, n, words, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	words := []string{
		"problem",
		"poem",
		"problem",
		"problem",
	}
	expect := int64(50)

	runSample(t, n, words, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	words := []string{
		"contest",
		"coolest",
		"unused",
	}
	expect := int64(4)

	runSample(t, n, words, expect)
}
