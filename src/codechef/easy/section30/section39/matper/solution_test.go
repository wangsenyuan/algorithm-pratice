package main

import "testing"

func runSample(t *testing.T, S string, P []string, expect int) {
	res := int(solve(S, P))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "abbabacbac"
	P := []string{
		"ab",
		"ac",
		"b",
	}
	expect := 3
	runSample(t, S, P, expect)
}

func TestSample2(t *testing.T) {
	S := "aaa"
	P := []string{
		"aa",
		"aa",
	}
	expect := 0
	runSample(t, S, P, expect)
}

func TestSample3(t *testing.T) {
	S := "aaaa"
	P := []string{
		"aa",
		"aa",
	}
	expect := 2
	runSample(t, S, P, expect)
}
