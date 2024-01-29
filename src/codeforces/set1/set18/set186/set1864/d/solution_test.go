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
		"00100",
		"01110",
		"11111",
		"11111",
		"11111",
	}
	expect := 1
	runSample(t, mat, expect)
}
