package main

import "testing"

func runSample(t *testing.T, n int, s string, expect string) {
	res := solve(n, s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s := "acba"
	expect := "abca"
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	s := "bcbaaabb"
	expect := "aabbbcba"
	runSample(t, n, s, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	s := "bdbcbccdbdbaaccd"
	expect := "abdbdccacbdbdccb"
	runSample(t, n, s, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	s := "ccfcffccccccffcfcfccfffffcccccff"
	expect := "cccccffffcccccffccfcffcccccfffff"
	runSample(t, n, s, expect)
}
