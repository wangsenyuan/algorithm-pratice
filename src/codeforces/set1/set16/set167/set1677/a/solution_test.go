package main

import "testing"

func runSample(t *testing.T, P []int, expect int) {
	res := solve(P)

	if int(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{5, 3, 6, 1, 4, 2}
	expect := 3
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{1, 2, 3, 4}
	expect := 0
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{5, 1, 6, 2, 8, 3, 4, 10, 9, 7}
	expect := 28
	runSample(t, P, expect)
}
