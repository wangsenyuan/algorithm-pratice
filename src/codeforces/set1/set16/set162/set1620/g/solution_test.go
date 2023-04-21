package main

import "testing"

func runSample(t *testing.T, S []string, expect int64) {
	res := solve(S)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := []string{
		"a", "b", "c",
	}
	var expect int64 = 92
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := []string{
		"aa", "a",
	}
	var expect int64 = 21
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := []string{
		"a", "a",
	}
	var expect int64 = 10
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := []string{
		"abcd", "aabb",
	}
	var expect int64 = 124
	runSample(t, S, expect)
}

func TestSample5(t *testing.T) {
	S := []string{
		"ddd",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaabbbbbbbbbbbcccccccccccciiiiiiiiiiiiiiiiiiiiiiooooooooooqqqqqqqqqqqqqqqqqqvvvvvzzzzzzzzzzzz",
	}
	var expect int64 = 15706243380
	runSample(t, S, expect)
}
