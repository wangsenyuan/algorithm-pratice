package main

import "testing"

func runSample(t *testing.T, a []int, s int, expect int) {
	res := solve(a, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3}
	s := 3
	expect := 2
	runSample(t, a, s, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2}
	s := 4
	expect := 1
	runSample(t, a, s, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3, 2}
	s := 5
	expect := 3
	runSample(t, a, s, expect)
}
