package main

import "testing"

func runSample(t *testing.T, B []int, expect int) {
	res := solve(len(B), B)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := []int{1, -1, 1}
	expect := 3
	runSample(t, B, expect)
}

func TestSample2(t *testing.T) {
	B := []int{2, -1, 1, -2, 2, 3, -5, 0, 2, -1}
	expect := 223
	runSample(t, B, expect)
}
