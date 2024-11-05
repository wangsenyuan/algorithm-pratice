package main

import "testing"

func runSample(t *testing.T, statements []string, expect int) {
	res := solve(statements)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	statements := []string{
		"add",
		"for 43",
		"end",
		"for 10",
		"for 15",
		"add",
		"end",
		"add",
		"end",
	}
	expect := 161
	runSample(t, statements, expect)
}

func TestSample2(t *testing.T) {
	statements := []string{
		"for 62",
		"end",
	}
	expect := 0
	runSample(t, statements, expect)
}

func TestSample3(t *testing.T) {
	statements := []string{
		"for 100",
		"for 100",
		"for 100",
		"for 100",
		"for 100",
		"add",
		"end",
		"end",
		"end",
		"end",
		"end",
	}
	expect := -1
	runSample(t, statements, expect)
}
