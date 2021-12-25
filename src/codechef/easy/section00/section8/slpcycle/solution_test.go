package main

import "testing"

func runSample(t *testing.T, n int, H int, S string, expect bool) {
	res := solve(n, H, S)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, H := 3, 1
	S := "111"
	expect := false
	runSample(t, n, H, S, expect)
}

func TestSample2(t *testing.T) {
	n, H := 7, 5
	S := "0000100"
	expect := true
	runSample(t, n, H, S, expect)
}

func TestSample3(t *testing.T) {
	n, H := 6, 5
	S := "000010"
	expect := false
	runSample(t, n, H, S, expect)
}
