package main

import "testing"

func runSample(t *testing.T, s string, a, b int) {
	c, d := solve(len(s), s)

	if a != c || b != d {
		t.Errorf("Sample expect %d, %d, but got %d %d", a, b, c, d)
	}
}

func TestSample1(t *testing.T) {
	s := "(())"
	a, b := 0, 1
	runSample(t, s, a, b)
}

func TestSample2(t *testing.T) {
	s := ")()("
	a, b := 2, 1
	runSample(t, s, a, b)
}

func TestSample3(t *testing.T) {
	s := "()(()"
	a, b := 1, 2
	runSample(t, s, a, b)
}

func TestSample4(t *testing.T) {
	s := ")(()())((("
	a, b := 4, 1
	runSample(t, s, a, b)
}