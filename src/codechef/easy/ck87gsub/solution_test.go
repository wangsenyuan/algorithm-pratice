package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve([]byte(s))

	if res != expect {
		t.Errorf("sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abba"
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := "a"
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := "ab"
	runSample(t, s, 0)
}
