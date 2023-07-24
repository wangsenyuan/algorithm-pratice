package main

import "testing"

func runSample(t *testing.T, c []int, k int, expect int) {
	res := solve(c, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	c := []int{1, 1, 2, 2, 1, 1, 2, 2, 2, 1}
	expect := 3
	runSample(t, c, k, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	c := []int{1, 2, 3, 4, 5, 6, 7}
	expect := 6
	runSample(t, c, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	c := []int{1, 3, 3, 3, 3, 1, 2, 1, 3, 3}
	expect := 2
	runSample(t, c, k, expect)
}
