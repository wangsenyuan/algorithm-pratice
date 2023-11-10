package main

import "testing"

func runSample(t *testing.T, w int, h int, a []int, b []int, expect bool) {
	res := solve(w, h, a, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w, h := 3, 2
	a := []int{13, 22, 29}
	b := []int{5, 16, 25}
	expect := false
	runSample(t, w, h, a, b, expect)
}

func TestSample2(t *testing.T) {
	w, h := 10, 5
	a := []int{65, 95, 165}
	b := []int{40, 65, 145}
	expect := true
	runSample(t, w, h, a, b, expect)
}
