package main

import "testing"

func runSample(t *testing.T, a int, b int, c []int, expect int) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 2, 7
	c := []int{3, 5, 12, 13, 21}
	expect := 173
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a, b := 27182, 31415
	c := []int{16, 18, 33, 98, 874, 989, 4848, 20458, 34365, 38117, 72030}
	expect := 3298918744
	runSample(t, a, b, c, expect)
}
