package main

import (
	"testing"
)

func runSample(t *testing.T, k int, s string, a, b int) {
	c, d := solve(k, s)

	if a != c || b != d {
		t.Errorf("Sample expect %d/%d, but got %d/%d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	k := 6
	s := "abcdefghij"
	a, b := 10, 1
	runSample(t, k, s, a, b)
}

func TestSample2(t *testing.T) {
	k := 1
	s := "qwertyuiop"
	a, b := 159, 1
	runSample(t, k, s, a, b)
}
