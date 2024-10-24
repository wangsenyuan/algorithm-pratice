package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 1, 2, 3}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 2, 3, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{5, 14, 4, 10, 2}
	expect := 3
	runSample(t, a, expect)
}
