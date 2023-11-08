package main

import "testing"

func runSample(t *testing.T, sheets []string, expect int) {
	res := solve(len(sheets), sheets)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	sheets := []string{
		"inn",
		"nww",
		"wii",
	}
	expect := 2
	runSample(t, sheets, expect)
}
