package main

import "testing"

func runSample(t *testing.T, r []int, expect int) {
	res := solve(r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := []int{2, 1}
	expect := 2
	runSample(t, r, expect)
}

func TestSample2(t *testing.T) {
	r := []int{2, 3, 4, 5, 1}
	expect := 2
	runSample(t, r, expect)
}

func TestSample3(t *testing.T) {
	r := []int{2, 1, 4, 2, 3}
	expect := 5
	runSample(t, r, expect)
}

func TestSample4(t *testing.T) {
	r := []int{4, 1, 1, 5, 4}
	expect := 4
	runSample(t, r, expect)
}

func TestSample5(t *testing.T) {
	r := []int{4, 3, 9, 1, 6, 7, 9, 10, 10, 3}
	expect := 5
	runSample(t, r, expect)
}
