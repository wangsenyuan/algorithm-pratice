package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 4}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 10, 6}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 8, 7, 2}
	expect := 8
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{9, 8, 2, 4, 6}
	expect := 16
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 6, 13, 3, 15, 5, 10, 8, 16, 9, 11, 18}
	expect := 232
	runSample(t, a, expect)
}
