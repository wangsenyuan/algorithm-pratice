package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 1, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 1, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 1, 5, 2, 4}
	expect := 11
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 6}
	expect := 5
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{6, 3, 4, 4, 4}
	expect := 13
	runSample(t, a, expect)
}
