package main

import "testing"

func runSample(t *testing.T, h int, a []int, c []int, expect int) {
	res := solve(h, a, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := 3
	a := []int{2, 1}
	c := []int{2, 1}
	expect := 1
	runSample(t, h, a, c, expect)
}

func TestSample2(t *testing.T) {
	h := 21
	a := []int{1, 1, 1, 1, 1, 1}
	c := []int{5, 5, 8, 10, 7, 6}
	expect := 21
	runSample(t, h, a, c, expect)
}

func TestSample3(t *testing.T) {
	h := 90000
	a := []int{200000, 200000}
	c := []int{1, 1}
	expect := 1
	runSample(t, h, a, c, expect)
}
