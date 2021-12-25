package main

import "testing"

func runSample(t *testing.T, S string, X string, U, V []int, expect int) {
	res := solve(len(X), len(S), len(U), S, X, U, V)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "aac"
	X := "aaca"
	U := []int{1, 2, 2, 1}
	V := []int{2, 3, 4, 2}
	expect := 3
	runSample(t, S, X, U, V, expect)
}

func TestSample2(t *testing.T) {
	S := "aa"
	X := "aa"
	U := []int{1}
	V := []int{2}
	expect := 1
	runSample(t, S, X, U, V, expect)
}
