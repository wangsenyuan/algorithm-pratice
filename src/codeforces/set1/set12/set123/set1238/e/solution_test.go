package main

import "testing"

func runSample(t *testing.T, s string, m int, expect int) {
	res := int(solve(s, m))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aacabc"
	m := 3
	expect := 5
	runSample(t, s, m, expect)
}

func TestSample2(t *testing.T) {
	s := "abacabadabacaba"
	m := 4
	expect := 16
	runSample(t, s, m, expect)
}
