package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 2}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 6, 9, 8, 55, 3, 14, 2, 7, 2}
	expect := false
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 11, 4, 1, 5, 1, 2}
	expect := true
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 6, 1, 5, 7, 8}
	expect := false
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 5, 10, 4, 4, 9, 6, 7, 8}
	expect := true
	runSample(t, a, expect)
}
