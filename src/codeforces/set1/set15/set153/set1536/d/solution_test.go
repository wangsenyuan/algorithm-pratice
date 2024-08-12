package main

import "testing"

func runSample(t *testing.T, b []int, expect bool) {
	res := solve(b)

	if expect != res {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	b := []int{6, 2, 1, 3}
	expect := false
	runSample(t, b, expect)
}

func TestSample2(t *testing.T) {
	b := []int{4, -8, 5, 6, -7}
	expect := false
	runSample(t, b, expect)
}

func TestSample3(t *testing.T) {
	b := []int{2, 1, 2, 3}
	expect := true
	runSample(t, b, expect)
}

func TestSample4(t *testing.T) {
	b := []int{-8, 2, -6, -5, -4, 3, 3, 2}
	expect := false
	runSample(t, b, expect)
}

func TestSample5(t *testing.T) {
	b := []int{1, 1, 3, 1, 0, -2, -1}
	expect := true
	runSample(t, b, expect)
}

func TestSample6(t *testing.T) {
	b := []int{6, 12, 8, 6, 2, 6, 10}
	expect := false
	runSample(t, b, expect)
}

func TestSample7(t *testing.T) {
	b := []int{5, 1, 2, 3, 6, 7}
	expect := false
	runSample(t, b, expect)
}
