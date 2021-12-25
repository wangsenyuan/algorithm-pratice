package main

import "testing"

func TestSample1(t *testing.T) {
	a, b, v := 0, 0, 0
	res := solve(a, b, v)
	if res != 1 {
		t.Errorf("sample: %d %d %d, should give 1, but got %d", a, b, v, res)
	}
}

func TestSample2(t *testing.T) {
	a, b, v := 1, 0, 1
	res := solve(a, b, v)
	if res != 1 {
		t.Errorf("sample: %d %d %d, should give 1, but got %d", a, b, v, res)
	}
}

func TestSample3(t *testing.T) {
	a, b, v := 3, 2, 6
	res := solve(a, b, v)
	if res != 3 {
		t.Errorf("sample: %d %d %d, should give 3, but got %d", a, b, v, res)
	}
}

func TestSample4(t *testing.T) {
	a, b, v := 4, 2, 8
	res := solve(a, b, v)
	if res != 4 {
		t.Errorf("sample: %d %d %d, should give 4, but got %d", a, b, v, res)
	}
}
