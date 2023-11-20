package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-1, -1, -1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, -5, 0, 2}
	expect := 13
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 6
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-1, 10, 9, 8, 7, 6}
	expect := 39
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 1, 8, -5, 1, -5, 8}
	expect := 30
	runSample(t, a, expect)
}
