package main

import (
	"index/suffixarray"
	"testing"
)

func runSample(t *testing.T, s string, b []string, expect int) {
	res := solve(s, b)

	if res[0] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	x := s[res[1] : res[1]+res[0]]

	sa := suffixarray.New([]byte(x))

	for _, w := range b {
		tmp := sa.Lookup([]byte(w), 1)
		if len(tmp) == 1 {
			t.Fatalf("Sample result %v, got substring %s, not correct", res, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "Go_straight_along_this_street"
	b := []string{
		"str",
		"long",
		"tree",
		"biginteger",
		"ellipse",
	}
	expect := 12
	runSample(t, s, b, expect)
}

func TestSample2(t *testing.T) {
	s := "unagioisii"
	b := []string{
		"ioi",
		"unagi",
	}
	expect := 5
	runSample(t, s, b, expect)
}

func TestSample3(t *testing.T) {
	s := "Z"
	b := []string{
		"a",
	}
	expect := 1
	runSample(t, s, b, expect)
}
