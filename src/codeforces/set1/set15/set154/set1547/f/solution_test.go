package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{16, 24, 10, 5}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{42, 42, 42, 42}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 6, 4}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 1
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{9, 9, 27, 9, 9, 63}
	expect := 1
	runSample(t, a, expect)
}
