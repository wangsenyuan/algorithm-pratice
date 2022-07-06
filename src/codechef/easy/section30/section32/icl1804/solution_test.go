package main

import "testing"

func runSample(t *testing.T, P int, C []int, D []int, expect int) {
	res := solve(P, C, D)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	D := []int{10, 8}
	C := []int{9, 1, 9, 1}
	P := 90
	expect := -1
	runSample(t, P, C, D, expect)
}

func TestSample2(t *testing.T) {
	D := []int{2, 3}
	C := []int{4, 4, 2, 2}
	P := 20
	expect := 7
	runSample(t, P, C, D, expect)
}

func TestSample3(t *testing.T) {
	D := []int{346, 36}
	C := []int{580, 403, 656, 10}
	P := 906997
	expect := 2092
	runSample(t, P, C, D, expect)
}
