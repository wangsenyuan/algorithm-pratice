package main

import (
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := solve(s)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	x := res[0]

	for i := 1; i < len(res); i++ {
		if res[i-1] >= res[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		x += res[i]
	}

	if x != s {
		t.Fatalf("Sample result %v, not correct", res)
	}

}

func TestSample1(t *testing.T) {
	s := "ABACUS"
	expect := []string{
		"A",
		"BA",
		"C",
		"US",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "AAAAAA"
	expect := []string{
		"A",
		"AA",
		"AAA",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "EDCBA"
	expect := []string{
		"EDCBA",
	}
	runSample(t, s, expect)
}
