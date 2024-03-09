package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 8, 2, 6, 4}
	expect := 16
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 7, 1, 3, 5}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 6}
	expect := 386
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 1, 1, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 1, 2, 1, 1, 2,
		2, 1, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 2, 2, 1, 1, 1, 1,
		2, 2, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 2, 2, 1, 2, 1, 1}
	expect := 0
	runSample(t, a, expect)
}
