package main

import "testing"

func runSample(t *testing.T, n, k int, S string, expect int) {
	res := solve(n, k, S)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 1
	S := "aa"
	expect := 5
	runSample(t, n, k, S, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 2
	S := "abcde"
	expect := 20
	runSample(t, n, k, S, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 2
	S := "abac"
	expect := 11
	runSample(t, n, k, S, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 1
	S := "abccbaabca"
	expect := 286
	runSample(t, n, k, S, expect)
}
