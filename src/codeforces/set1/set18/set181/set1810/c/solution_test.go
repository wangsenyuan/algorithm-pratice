package main

import "testing"

func runSample(t *testing.T, a []int, c int, d int, expect int) {
	res := solve(a, c, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c, d := 1, 5
	a := []int{1, 2, 3, 5, 6}
	expect := 2
	runSample(t, a, c, d, expect)
}

func TestSample2(t *testing.T) {
	c, d := 1000000000, 1
	a := []int{1000000000, 1}
	expect := 999999998
	runSample(t, a, c, d, expect)
}

func TestSample3(t *testing.T) {
	c, d := 2, 3
	a := []int{1, 1, 1, 3, 3}
	expect := 8
	runSample(t, a, c, d, expect)
}

func TestSample4(t *testing.T) {
	c, d := 1, 1
	a := []int{2, 2}
	expect := 2
	runSample(t, a, c, d, expect)
}
