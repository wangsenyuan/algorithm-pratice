package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaza"
	var expect int64 = 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "cbaabc"
	var expect int64 = 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "icpcsguru"
	var expect int64 = 30
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "fhakqhdhrgfjruxndgfhdvcxhsrjfgdhsyrhfjcbfgdvrtdysf"
	var expect int64 = 482
	runSample(t, s, expect)
}
