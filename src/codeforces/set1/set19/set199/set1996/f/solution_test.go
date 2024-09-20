package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 6, 7}
	b := []int{2, 3, 4}
	k := 4
	expect := 21
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{32, 52, 68, 64, 14}
	b := []int{18, 14, 53, 24, 8}
	k := 9
	expect := 349
	runSample(t, a, b, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 4, 3, 2, 1}
	k := 1000
	expect := 27
	runSample(t, a, b, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1000000}
	b := []int{1}
	k := 1000000
	expect := 500000500000
	runSample(t, a, b, k, expect)
}

func TestSample5(t *testing.T) {
	a := []int{3, 3, 5, 10, 6, 8, 6, 8, 7, 7}
	b := []int{6, 1, 7, 4, 1, 1, 8, 9, 3, 1}
	k := 6
	expect := 47
	runSample(t, a, b, k, expect)
}
