package main

import "testing"

func runSample(t *testing.T, mat []string, expect int) {
	res := solve(len(mat), mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		"E7",
		"E7",
		"E7",
		"00",
		"00",
		"E7",
		"E7",
		"E7",
	}
	expect := 1
	runSample(t, mat, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		"FF",
		"FF",
		"FF",
		"FF",
		"FF",
		"FF",
		"FF",
		"FF",
	}
	expect := solve1(len(mat), mat)

	runSample(t, mat, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"C3",
		"C3",
		"C3",
		"C3",
		"C3",
		"C3",
		"C3",
		"C3",
	}
	expect := solve1(len(mat), mat)
	runSample(t, mat, expect)
}

