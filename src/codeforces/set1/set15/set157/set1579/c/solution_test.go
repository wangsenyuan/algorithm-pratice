package main

import "testing"

func runSample(t *testing.T, mat []string, k int, expect bool) {
	res := solve(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	mat := []string{
		"*.*",
		"...",
	}
	expect := false
	runSample(t, mat, k, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	mat := []string{
		"*.*.*...*",
		".*.*...*.",
		"..*.*.*..",
		".....*...",
	}
	expect := true
	runSample(t, mat, k, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	mat := []string{
		"*.*.",
		"****",
		".**.",
		"....",
	}
	expect := true
	runSample(t, mat, k, expect)
}
