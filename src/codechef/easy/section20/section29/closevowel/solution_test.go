package main

import "testing"

func runSample(t *testing.T, n int, s string, expect int) {
	res := solve(n, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	s := "aeiou"
	expect := 1
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	s := "abcde"
	expect := 2
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	s := "codechef"
	expect := 4
	runSample(t, n, s, expect)
}
