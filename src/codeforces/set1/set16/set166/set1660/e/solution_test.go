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
		"010",
		"011",
		"100",
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"10",
		"10",
	}
	expect := 2
	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"1111",
		"1011",
		"1111",
		"1111",
	}
	expect := 11
	runSample(t, mat, expect)
}
