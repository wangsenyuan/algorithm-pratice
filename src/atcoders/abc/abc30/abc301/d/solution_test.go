package main

import "testing"

func runSample(t *testing.T, s string, n int64, expect int64) {
	res := solve(s, n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "?0?"
	var n int64 = 2
	var expect int64 = 1
	runSample(t, s, n, expect)
}

func TestSample2(t *testing.T) {
	s := "101"
	var n int64 = 4
	var expect int64 = -1
	runSample(t, s, n, expect)
}

func TestSample3(t *testing.T) {
	s := "?0?"
	var n int64 = 1000000000000000000
	var expect int64 = 5
	runSample(t, s, n, expect)
}

func TestSample4(t *testing.T) {
	s := "?0?"
	var n int64 = 4
	var expect int64 = 4
	runSample(t, s, n, expect)
}
