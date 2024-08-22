package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	k := 2
	expect := "acac"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abc"
	k := 1
	expect := "abc"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "abaabaaaa"
	k := 3
	expect := "abaabaaab"
	runSample(t, s, k, expect)
}
