package main

import "testing"

func runSample(t *testing.T, p []int, c []int, expect int) {
	res := solve(p, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{1, 3, 4, 2}
	c := []int{1, 2, 2, 3}
	expect := 1
	runSample(t, p, c, expect)
}

func TestSample2(t *testing.T) {
	p := []int{2, 3, 4, 5, 1}
	c := []int{1, 2, 3, 4, 5}
	expect := 5
	runSample(t, p, c, expect)
}

func TestSample3(t *testing.T) {
	p := []int{7, 4, 5, 6, 1, 8, 3, 2}
	c := []int{5, 3, 6, 4, 7, 5, 8, 4}
	expect := 2
	runSample(t, p, c, expect)
}
