package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect bool) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 2, 4}
	b := []int{1, 3, 3, 2, 4}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 2, 2, 4}
	b := []int{3, 4, 3, 4, 4}
	expect := false
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 1, 1, 1}
	b := []int{3, 3, 3, 2, 2}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 1}
	b := []int{1, 2}
	expect := false
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 1, 2}
	b := []int{2, 1, 2}
	expect := false
	runSample(t, a, b, expect)
}
