package main

import "testing"

func runSample(t *testing.T, n int, S []string, expect int64) {
	res := solve1(n, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []string{
		"bbcab",
		"aa",
		"aab",
	}
	var expect int64 = 4
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	S := []string{
		"bbc",
		"bcaa",
		"abb",
	}
	var expect int64 = 4
	runSample(t, n, S, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	S := []string{
		"aa",
		"acc",
		"abbca",
	}
	var expect int64 = 2
	runSample(t, n, S, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	S := []string{
		"aa",
		"aabb",
		"bb",
	}
	var expect int64 = 4
	runSample(t, n, S, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	S := []string{
		"cab",
		"aabc",
		"bbc",
	}
	var expect int64 = 4
	runSample(t, n, S, expect)
}

func TestSample6(t *testing.T) {
	n := 3
	S := []string{
		"a",
		"b",
		"c",
	}
	var expect int64 = 0
	runSample(t, n, S, expect)
}
