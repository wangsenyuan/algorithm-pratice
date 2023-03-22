package main

import "testing"

func runSample(t *testing.T, s string, k int, x int64, expect string) {
	res := solve(s, k, x)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "a*"
	k := 4
	var x int64 = 3
	expect := "abb"
	runSample(t, s, k, x, expect)
}

func TestSample2(t *testing.T) {
	s := "a**a"
	k := 1
	var x int64 = 3
	expect := "abba"
	runSample(t, s, k, x, expect)
}

func TestSample3(t *testing.T) {
	s := "**a***"
	k := 3
	var x int64 = 20
	expect := "babbbbbbbbb"
	runSample(t, s, k, x, expect)
}

// 2 0 1
//a*

func TestSample4(t *testing.T) {
	s := "a*"
	k := 0
	var x int64 = 1
	expect := "a"
	runSample(t, s, k, x, expect)
}
