package main

import "testing"

func runSample(t *testing.T, s string, r string, expect int64) {
	res := solve(s, r)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a"
	r := "a"
	var expect int64 = -1
	runSample(t, s, r, expect)
}
func TestSample2(t *testing.T) {
	s := "rll"
	r := "rrr"
	var expect int64 = 0
	runSample(t, s, r, expect)
}

func TestSample3(t *testing.T) {
	s := "caa"
	r := "aca"
	var expect int64 = 2
	runSample(t, s, r, expect)
}

func TestSample4(t *testing.T) {
	s := "ababa"
	r := "aabba"
	var expect int64 = 2
	runSample(t, s, r, expect)
}
