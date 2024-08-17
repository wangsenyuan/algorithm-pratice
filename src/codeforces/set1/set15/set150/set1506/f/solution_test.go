package main

import "testing"

func runSample(t *testing.T, r []int, c []int, expect int) {
	res := solve(r, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := []int{1, 4, 2}
	c := []int{1, 3, 1}
	expect := 0
	runSample(t, r, c, expect)
}

func TestSample2(t *testing.T) {
	r := []int{2, 4}
	c := []int{2, 3}
	expect := 1
	runSample(t, r, c, expect)
}

func TestSample3(t *testing.T) {
	r := []int{1, 1000000000}
	c := []int{1, 1000000000}
	expect := 999999999
	runSample(t, r, c, expect)
}

func TestSample4(t *testing.T) {
	r := []int{3, 10, 5, 8}
	c := []int{2, 5, 2, 4}
	expect := 2
	runSample(t, r, c, expect)
}
