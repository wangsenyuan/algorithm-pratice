package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "bbab"
	k := 2
	expect := "aaaa"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "cceddda"
	k := 5
	expect := "baccacd"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "ecdaed"
	k := 5
	expect := "aabdac"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "ccabbaca"
	k := 3
	expect := "aaaaaaaa"
	runSample(t, s, k, expect)
}
