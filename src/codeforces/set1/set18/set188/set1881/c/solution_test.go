package main

import "testing"

func runSample(t *testing.T, mat []string, expect int) {
	res := solve(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"abba",
		"bcbb",
		"bccb",
		"abba",
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"codefo",
		"rcesco",
		"deforc",
		"escode",
		"forces",
		"codefo",
	}
	expect := 181
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"bbaa",
		"abba",
		"aaba",
		"abba",
	}
	expect := 9
	runSample(t, mat, expect)
}
