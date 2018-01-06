package main

import "testing"

func TestSample1(t *testing.T) {
	L, R := 1, 5
	expect := 6
	res := solve(L, R)
	if res != expect {
		t.Errorf("sample: %d %d, should give %d, but got %d", L, R, expect, res)
	}
}

func TestSample2(t *testing.T) {
	L, R := 6, 10
	expect := 30
	res := solve(L, R)
	if res != expect {
		t.Errorf("sample: %d %d, should give %d, but got %d", L, R, expect, res)
	}
}
