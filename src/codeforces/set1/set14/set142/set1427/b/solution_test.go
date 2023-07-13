package main

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "WLLWLW"
	k := 1
	expect := 6
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "LLWLWLWWWLWLLWLWWWLWLLWLLWLLLLWLLWWWLWWL"
	k := 7
	expect := 46
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "LLLL"
	k := 3
	expect := 5
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "WLLWL"
	k := 2
	expect := 7
	runSample(t, s, k, expect)
}
