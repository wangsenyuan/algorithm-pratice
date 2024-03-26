package main

import "testing"

func runSample(t *testing.T, s string, expect []int) {
	even, odd := solve(s)

	if even != expect[0] || odd != expect[1] {
		t.Fatalf("Sample expect %v, but got %d %d", expect, even, odd)
	}
}

func TestSample1(t *testing.T) {
	s := "bb"
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "baab"
	expect := []int{2, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "babb"
	expect := []int{2, 5}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "babaa"
	expect := []int{2, 7}
	runSample(t, s, expect)
}
