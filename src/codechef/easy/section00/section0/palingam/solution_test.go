package main

import "testing"

func runSample(t *testing.T, a, b string, expect bool) {
	res := solve(a, b)
	if res != expect {
		t.Errorf("sample %s %s, expect %t, but got %t", a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := "ab", "ba"
	runSample(t, a, b, false)
}

func TestSample2(t *testing.T) {
	a, b := "aba", "cde"
	runSample(t, a, b, true)
}

func TestSample3(t *testing.T) {
	a, b := "ab", "cd"
	runSample(t, a, b, false)
}
