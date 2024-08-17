package main

import "testing"

func runSample(t *testing.T, a []int, b []int, c int, expect int) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 10, 11}
	b := []int{1, 2, 7}
	c := 15
	expect := 6
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, 10, 50}
	b := []int{3, 14, 12}
	c := 100
	expect := 13
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1}
	b := []int{1}
	c := 1000000000
	expect := 1000000000
	runSample(t, a, b, c, expect)
}
