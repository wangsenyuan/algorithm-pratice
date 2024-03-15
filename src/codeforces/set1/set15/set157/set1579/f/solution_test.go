package main

import "testing"

func runSample(t *testing.T, a []int, d int, expect int) {
	res := solve(a, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1}
	d := 1
	expect := 1
	runSample(t, a, d, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 0}
	d := 2
	expect := 1
	runSample(t, a, d, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 0, 1, 0}
	d := 2
	expect := 3
	runSample(t, a, d, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 1, 0, 1}
	d := 2
	expect := -1
	runSample(t, a, d, expect)
}

func TestSample5(t *testing.T) {
	a := []int{0}
	d := 1
	expect := 0
	runSample(t, a, d, expect)
}
