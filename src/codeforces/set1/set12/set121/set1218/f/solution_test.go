package main

import "testing"

func runSample(t *testing.T, k int, x []int, a int, c []int, expect int) {
	res := solve(k, x, a, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 10000
	x := []int{10000, 30000, 30000, 40000, 20000}
	a := 20000
	c := []int{5, 2, 8, 3, 6}
	expect := 5
	runSample(t, k, x, a, c, expect)
}

func TestSample2(t *testing.T) {
	k := 10000
	x := []int{10000, 40000, 30000, 40000, 20000}
	a := 10000
	c := []int{5, 2, 8, 3, 6}
	expect := -1
	runSample(t, k, x, a, c, expect)
}
