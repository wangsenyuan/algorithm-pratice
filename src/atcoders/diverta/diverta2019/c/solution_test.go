package main

import "testing"

func runSample(t *testing.T, s []string, expect int) {
	res := solve(s)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []string{
		"ABCA",
		"XBAZ",
		"BAD",
	}
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []string{
		"BEWPVCRWH",
		"ZZNQYIJX",
		"BAVREA",
		"PA",
		"HJMYITEOX",
		"BCJHMRMNK",
		"BP",
		"QVFABZ",
		"PRGKSPUNA",
	}
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []string{
		"RABYBBE",
		"JOZ",
		"BMHQUVA",
		"BPA",
		"ISU",
		"MCMABAOBHZ",
		"SZMEHMA",
	}
	expect := 4
	runSample(t, s, expect)
}
