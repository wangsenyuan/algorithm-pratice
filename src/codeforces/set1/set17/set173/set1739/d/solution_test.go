package main

import "testing"

func runSample(t *testing.T, n int, k int, P []int, expect int) {
	res := solve(n, k, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	k := 1
	P := []int{1, 1, 2, 2}
	expect := 2
	runSample(t, n, k, P, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	k := 2
	P := []int{1, 1, 2, 2}
	expect := 1
	runSample(t, n, k, P, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	k := 0
	P := []int{1, 2, 3, 4, 5}
	expect := 5
	runSample(t, n, k, P, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	k := 1
	P := []int{1, 2, 3, 4, 5}
	expect := 3
	runSample(t, n, k, P, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	k := 3
	P := []int{1, 1, 1}
	expect := 1
	runSample(t, n, k, P, expect)
}
