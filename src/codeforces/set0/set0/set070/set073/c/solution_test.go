package main

import "testing"

func runSample(t *testing.T, s string, k int, bonus []string, expect int) {
	res := solve(s, k, bonus)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "winner"
	k := 4
	bonus := []string{
		"s e 7",
		"o s 8",
		"l o 13",
		"o o 8",
	}
	expect := 36
	runSample(t, s, k, bonus, expect)
}

func TestSample2(t *testing.T) {
	s := "abcdef"
	k := 1
	bonus := []string{
		"a b -10",
		"b c 5",
		"c d 5",
		"d e 5",
		"e f 5",
	}
	expect := 20
	runSample(t, s, k, bonus, expect)
}
