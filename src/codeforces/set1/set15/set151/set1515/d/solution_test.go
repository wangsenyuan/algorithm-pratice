package main

import "testing"

func runSample(t *testing.T, n int, l int, r int, c []int, expect int) {
	res := solve(n, l, r, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, l, r := 6, 3, 3
	c := []int{1, 2, 3, 2, 2, 2}
	expect := 2
	runSample(t, n, l, r, c, expect)
}

func TestSample2(t *testing.T) {
	n, l, r := 6, 2, 4
	c := []int{1, 1, 2, 2, 2, 2}
	expect := 3
	runSample(t, n, l, r, c, expect)
}

func TestSample3(t *testing.T) {
	n, l, r := 6, 5, 1
	c := []int{6, 5, 4, 3, 2, 1}
	expect := 5
	runSample(t, n, l, r, c, expect)
}

func TestSample4(t *testing.T) {
	n, l, r := 4, 0, 4
	c := []int{4, 4, 4, 3}
	expect := 3
	runSample(t, n, l, r, c, expect)
}
