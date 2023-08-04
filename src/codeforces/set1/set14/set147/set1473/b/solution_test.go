package main

import "testing"

func runSample(t *testing.T, a string, b string, expect string) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "baba"
	b := "ba"
	expect := "baba"
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "aa"
	b := "aaa"
	expect := "aaaaaa"
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "aba"
	b := "ab"
	expect := "-1"
	runSample(t, a, b, expect)
}
