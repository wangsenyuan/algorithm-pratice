package main

import "testing"

func runSample(t *testing.T, S, T string, expect int) {
	res := solve(S, T)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "aa"
	T := "aa"
	expect := 5
	runSample(t, S, T, expect)
}

func TestSample2(t *testing.T) {
	S := "codeforces"
	T := "forceofcode"
	expect := 60
	runSample(t, S, T, expect)
}
