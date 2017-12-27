package main

import "testing"

func TestSample1(t *testing.T) {
	n, m, r, k := 1, 1, 1, 1
	res := solve(n, m, r, k)
	if res != 0 {
		t.Errorf("sample %d %d %d %d should give 0, but got %d", n, m, r, k, res)
	}
}

func TestSample2(t *testing.T) {
	n, m, r, k := 2, 2, 1, 1
	res := solve(n, m, r, k)
	if res != 2 {
		t.Errorf("sample %d %d %d %d should give 2, but got %d", n, m, r, k, res)
	}
}

func TestSample3(t *testing.T) {
	n, m, r, k := 2, 3, 1, 1
	res := solve(n, m, r, k)
	if res != 4 {
		t.Errorf("sample %d %d %d %d should give 4, but got %d", n, m, r, k, res)
	}
}
