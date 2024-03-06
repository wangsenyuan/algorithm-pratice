package main

import "testing"

func runSample(t *testing.T, n int, statements []string, expect int) {
	res := solve(n, statements)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	statements := []string{
		"1 2 imposter",
		"2 3 crewmate",
	}
	expect := 2
	runSample(t, n, statements, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	statements := []string{
		"1 3 crewmate",
		"2 5 crewmate",
		"2 4 imposter",
		"3 4 imposter",
	}
	expect := 4
	runSample(t, n, statements, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	statements := []string{
		"1 2 imposter",
		"2 1 crewmate",
	}
	expect := -1
	runSample(t, n, statements, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	statements := []string{
		"1 2 imposter",
		"1 2 imposter",
		"3 2 crewmate",
		"3 2 crewmate",
		"1 3 imposter",
	}
	expect := 2
	runSample(t, n, statements, expect)
}
