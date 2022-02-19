package main

import "testing"

func runSample(t *testing.T, B []int, expect bool) {
	res := solve(B)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{1, 1}
	expect := true
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{16, 2, 2022, 2001}
	expect := false
	runSample(t, B, expect)
}

func TestSample3(t *testing.T) {
	B := []int{2, 4, -2, 4}
	expect := true
	runSample(t, B, expect)
}
