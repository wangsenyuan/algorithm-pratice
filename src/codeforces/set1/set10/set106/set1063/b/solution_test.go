package main

import "testing"

func runSample(t *testing.T, mat []string, r int, c int, x int, y int, expect int) {
	res := solve(mat, r, c, x, y)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r, c := 3, 2
	x, y := 1, 2
	mat := []string{
		".....",
		".***.",
		"...**",
		"*....",
	}
	expect := 10
	runSample(t, mat, r, c, x, y, expect)
}

func TestSample2(t *testing.T) {
	r, c := 2, 2
	x, y := 0, 1
	mat := []string{
		"....",
		"..*.",
		"....",
		"....",
	}
	expect := 7
	runSample(t, mat, r, c, x, y, expect)
}
