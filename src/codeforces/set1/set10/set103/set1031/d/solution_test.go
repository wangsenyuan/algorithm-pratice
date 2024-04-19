package main

import "testing"

func runSample(t *testing.T, mat []string, k int, expect string) {
	res := solve(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"abcd",
		"bcde",
		"bcad",
		"bcde",
	}
	k := 2
	expect := "aaabcde"
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"bwwwz",
		"hrhdh",
		"sepsp",
		"sqfaf",
		"ajbvw",
	}
	k := 3
	expect := "aaaepfafw"
	runSample(t, mat, k, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"ypnxnnp",
		"pnxonpm",
		"nxanpou",
		"xnnpmud",
		"nhtdudu",
		"npmuduh",
		"pmutsnz",
	}
	k := 6
	expect := "aaaaaaadudsnz"
	runSample(t, mat, k, expect)
}
