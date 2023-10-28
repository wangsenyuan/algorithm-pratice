package main

import "testing"

func runSample(t *testing.T, p []int, a []int, d int, expect int) {
	res := solve(p, a, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1, 2, 3, 4}
	a := []int{1, 3}
	d := 2
	expect := 1
	runSample(t, p, a, d, expect)
}

func TestSample2(t *testing.T) {
	p := []int{5, 4, 3, 2, 1}
	a := []int{5, 2}
	d := 4
	expect := 3
	runSample(t, p, a, d, expect)
}

func TestSample3(t *testing.T) {
	p := []int{3, 4, 1, 5, 2}
	a := []int{3, 1, 2}
	d := 3
	expect := 2
	runSample(t, p, a, d, expect)
}

func TestSample4(t *testing.T) {
	p := []int{1, 2}
	a := []int{2, 1}
	d := 1
	expect := 0
	runSample(t, p, a, d, expect)
}

func TestSample5(t *testing.T) {
	p := []int{1, 2, 3, 4, 5, 6}
	a := []int{2, 5}
	d := 4
	expect := 2
	runSample(t, p, a, d, expect)
}
