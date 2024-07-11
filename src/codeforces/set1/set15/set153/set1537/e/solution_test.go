package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "dbcadabc"
	k := 16
	expect := "dbcadabcdbcadabc"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	k := 5
	expect := "aaaaa"
	runSample(t, s, k, expect)
}