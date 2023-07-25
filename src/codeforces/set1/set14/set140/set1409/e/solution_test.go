package main

import "testing"

func runSample(t *testing.T, points []int, k int, expect int) {
	res := solve(points, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{1, 5, 2, 3, 1, 5, 4}
	k := 1
	expect := 6
	runSample(t, x, k, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1000000000}
	k := 1
	expect := 1
	runSample(t, x, k, expect)
}

func TestSample3(t *testing.T) {
	x := []int{10, 7, 5, 15, 8}
	k := 10
	expect := 5
	runSample(t, x, k, expect)
}

func TestSample4(t *testing.T) {
	x := []int{15, 19, 8, 17, 20, 10, 9, 2, 10, 19}
	k := 10
	expect := 10
	runSample(t, x, k, expect)
}
