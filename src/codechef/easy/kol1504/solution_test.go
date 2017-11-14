package main

import "testing"

func TestSample1(t *testing.T) {
	n, d := 4, 2
	a, b := "qnyu", "ynqu"
	res := solve(n, d, a, b)
	if !res {
		t.Errorf("solve(%s, %s) should be true, but got false", a, b)
	}
}

func TestSample2(t *testing.T) {
	n, d := 4, 1
	a, b := "fbnc", "nbcf"
	res := solve(n, d, a, b)
	if !res {
		t.Errorf("solve(%s, %s) should be true, but got false", a, b)
	}
}

func TestSample3(t *testing.T) {
	n, d := 5, 2
	a, b := "abcde", "edacb"
	res := solve(n, d, a, b)
	if res {
		t.Errorf("solve(%s, %s) should be false, but got true", a, b)
	}
}

func TestSample4(t *testing.T) {
	n, d := 5, 2
	a, b := "abcde", "edabc"
	res := solve(n, d, a, b)
	if !res {
		t.Errorf("solve(%s, %s) should be true, but got false", a, b)
	}
}

func TestSample5(t *testing.T) {
	n, d := 5, 2
	a, b := "eff", "bae"
	res := solve(n, d, a, b)
	if res {
		t.Errorf("solve(%s, %s) should be false, but got true", a, b)
	}
}
