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
		".*.",
		"***",
	}
	expect := 5
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"..*.*..",
		".*****.",
		"*******",
		".*****.",
		"..*.*..",
	}
	expect := 34
	runSample(t, mat, expect)
}
