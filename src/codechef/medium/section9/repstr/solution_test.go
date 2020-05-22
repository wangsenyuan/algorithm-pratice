package main

import "testing"

func runSample(t *testing.T, L, H int, S string, a, b int) {
	c, d := solve(L, H, S)

	if a != c || b != d {
		t.Errorf("Sample %d %d %s, expect %d, %d, but got %d %d", L, H, S, a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	L, H := 3, 5
	S := "aabcbcbca"
	a, b := 2, 4
	runSample(t, L, H, S, a, b)
}

func TestSample2(t *testing.T) {
	L, H := 3, 5
	S := "baaaababababbababbab"
	a, b := 6, 3
	runSample(t, L, H, S, a, b)
}

func TestSample3(t *testing.T) {
	L, H := 1, 4
	S := "abcd"
	a, b := 1, 4
	runSample(t, L, H, S, a, b)
}
