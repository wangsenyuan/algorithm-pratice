package main

import "testing"

func runSample(t *testing.T, n int, pref []int, suf []int, expect int) {
	res := solve(n, pref, suf)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	pref := []int{1}
	suf := []int{1}
	expect := 1
	runSample(t, n, pref, suf, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	pref := []int{1, 2}
	suf := []int{2, 3, 4}
	expect := 3
	runSample(t, n, pref, suf, expect)
}

func TestSample3(t *testing.T) {
	n := 20
	pref := []int{1, 2, 3, 4, 12}
	suf := []int{12, 13, 18, 20}
	expect := 317580808
	runSample(t, n, pref, suf, expect)
}
