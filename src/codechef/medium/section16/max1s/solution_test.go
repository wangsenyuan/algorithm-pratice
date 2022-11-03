package main

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := solve(s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111"
	var expect int64 = 10
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "000"
	var expect int64 = 10
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "00100"
	var expect int64 = 26
	runSample(t, s, expect)
}
