package main

import "testing"

func runSample(t *testing.T, a []int, b []int, c []int, expect int) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{9, 6, 7, 5, 5}
	b := []int{8, 4, 5, 1, 2}
	c := []int{10, 4, 7}
	expect := 12
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 20, 20}
	b := []int{0, 0, 0}
	c := []int{9, 10, 19, 20}
	expect := 8
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 20, 20}
	b := []int{0, 0, 0}
	c := []int{9, 10, 19, 20}
	expect := 8
	runSample(t, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	a := []int{3}
	b := []int{1}
	c := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	expect := 4999999990
	runSample(t, a, b, c, expect)
}
