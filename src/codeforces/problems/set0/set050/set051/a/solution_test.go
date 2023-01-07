package main

import "testing"

func runSample(t *testing.T, amulets [][]string, expect int) {
	res := solve(amulets)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	amulets := [][]string{
		{"31", "23"},
		{"31", "23"},
		{"13", "32"},
		{"32", "13"},
	}

	expect := 1
	runSample(t, amulets, expect)

}
