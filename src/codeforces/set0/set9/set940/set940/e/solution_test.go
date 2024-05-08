package main

import "testing"

func runSample(t *testing.T, a []int, c int, expect int) {
	res := solve(a, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	c := 5
	expect := 6
	runSample(t, a, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 10, 10, 10, 10, 10, 10, 9, 10, 10, 10}
	c := 10
	expect := 92
	runSample(t, a, c, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 3, 6, 4, 5, 7, 1}
	c := 2
	expect := 17
	runSample(t, a, c, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 3, 4, 5, 5, 3, 4, 1}
	c := 4
	expect := 23
	runSample(t, a, c, expect)
}
