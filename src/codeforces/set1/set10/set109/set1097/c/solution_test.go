package main

import "testing"

func runSample(t *testing.T, words []string, expect int) {
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		")())",
		")",
		"((",
		"((",
		"(",
		")",
		")",
	}
	expect := 2
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"(",
		"((",
		"(((",
		"(())",
	}
	expect := 0
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"(())",
		"()",
	}
	expect := 1
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{
		")))",
		"())))",
		"(((",
		"((()",
		"()(",
		")((())",
		"(((((",
		")())(",
		")()((()",
		"((((()()(",
	}
	expect := 1
	runSample(t, words, expect)
}
