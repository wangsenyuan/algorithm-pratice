package main

import "testing"

func runSample(t *testing.T, l []int, r []int, c []int, expect int) {
	res := solve(l, r, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l := []int{8, 3}
	r := []int{12, 23}
	c := []int{100, 100}
	expect := 2400
	runSample(t, l, r, c, expect)
}

func TestSample2(t *testing.T) {
	l := []int{20, 1, 2, 5}
	r := []int{30, 4, 3, 10}
	c := []int{2, 3, 2, 3}
	expect := 42
	runSample(t, l, r, c, expect)
}
