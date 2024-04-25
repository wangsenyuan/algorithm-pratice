package main

import "testing"

func runSample(t *testing.T, a, b, c []int, expect int) {
	res := solve(a, b, c)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 0}
	b := []int{2, 0}
	c := []int{1, 1}
	expect := 4
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 0}
	b := []int{5, 0}
	c := []int{10, 0}
	expect := 11
	runSample(t, a, b, c, expect)
}
