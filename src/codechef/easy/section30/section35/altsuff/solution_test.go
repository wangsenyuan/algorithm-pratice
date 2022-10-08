package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	k := 1
	expect := "010"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "10001"
	k := 2
	expect := "10101"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "10011010111000"
	k := 3
	expect := "01100101010101"
	runSample(t, s, k, expect)
}
