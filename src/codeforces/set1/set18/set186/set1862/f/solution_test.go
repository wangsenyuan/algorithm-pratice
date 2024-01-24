package main

import "testing"

func runSample(t *testing.T, w int, f int, s []int, expect int) {
	res := solve(w, f, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	w, f := 2, 3
	s := []int{2, 6, 7}
	expect := 3
	runSample(t, w, f, s, expect)
}

func TestSample2(t *testing.T) {
	w, f := 37, 58
	s := []int{93}
	expect := 2
	runSample(t, w, f, s, expect)
}

func TestSample3(t *testing.T) {
	w, f := 37, 58
	s := []int{93}
	expect := 2
	runSample(t, w, f, s, expect)
}

func TestSample4(t *testing.T) {
	w, f := 190, 90
	s := []int{23, 97}
	expect := 1
	runSample(t, w, f, s, expect)
}

func TestSample5(t *testing.T) {
	w, f := 13, 4
	s := []int{10, 10, 2, 45}
	expect := 5
	runSample(t, w, f, s, expect)
}
