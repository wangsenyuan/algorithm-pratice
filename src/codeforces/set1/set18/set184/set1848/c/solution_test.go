package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect bool) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 0, 0, 0}
	b := []int{1, 2, 3, 4}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2}
	b := []int{2, 1}
	expect := false
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{100, 23, 53, 11, 56, 32}
	b := []int{1245, 31, 12, 6, 6, 6}
	expect := false
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{7, 6, 5, 4, 3, 2, 1}
	expect := true
	runSample(t, a, b, expect)
}
