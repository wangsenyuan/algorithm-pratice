package main

import "testing"

func runSample(t *testing.T, s string, m, p int, expect int) {
	res := solve([]byte(s), m, p)

	if res != expect {
		t.Errorf("Sample %s %d %d, expect %d, but got %d", s, m, p, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "AbcDef"
	m, p := 1, 2
	expect := 4
	runSample(t, s, m, p, expect)
}

func TestSample2(t *testing.T) {
	s := "abcabc"
	m, p := 1, 1
	expect := 3
	runSample(t, s, m, p, expect)
}
