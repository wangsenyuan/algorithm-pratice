package main

import "testing"

func runSample(t *testing.T, k int, h []int, p []int, expect bool) {
	res := solve(k, h, p)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 7
	h := []int{18, 5, 13, 9, 10, 1}
	p := []int{2, 7, 2, 1, 2, 6}
	expect := true
	runSample(t, k, h, p, expect)
}

func TestSample2(t *testing.T) {
	k := 4
	h := []int{5, 5, 5}
	p := []int{4, 4, 4}
	expect := false
	runSample(t, k, h, p, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	h := []int{2, 1, 3}
	p := []int{1, 1, 1}
	expect := true
	runSample(t, k, h, p, expect)
}
