package main

import "testing"

func runSample(t *testing.T, k int, start, end string, expect int) {
	res := solve(k, start, end)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	start := "ab"
	end := "ab"
	expect := 1
	runSample(t, k, start, end, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	start := "ababab"
	end := "ababab"
	expect := 2
	runSample(t, k, start, end, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	start := "ab"
	end := "ba"
	expect := 0
	runSample(t, k, start, end, expect)
}
