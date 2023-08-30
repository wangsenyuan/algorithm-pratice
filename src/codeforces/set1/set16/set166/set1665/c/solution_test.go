package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect int) {
	res := solve(n, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	P := []int{1, 1, 1, 2, 2, 4}
	expect := 4
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	P := []int{5, 5, 1, 4}
	expect := 4
	runSample(t, n, P, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	P := []int{1}
	expect := 2
	runSample(t, n, P, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	P := []int{3, 1}
	expect := 3
	runSample(t, n, P, expect)
}

func TestSample5(t *testing.T) {
	n := 6
	P := []int{1, 1, 1, 1, 1}
	expect := 4
	runSample(t, n, P, expect)
}
