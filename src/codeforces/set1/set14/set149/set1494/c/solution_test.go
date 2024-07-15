package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{-1, 1, 5, 11, 15}
	b := []int{-4, -3, -2, 6, 7, 15}
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 1}
	b := []int{-1000000000, 1000000000}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1000000000, 1000000000}
	b := []int{-1, 1}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0}
	b := []int{1}
	expect := 1
	runSample(t, a, b, expect)
}
