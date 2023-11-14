package main

import "testing"

func runSample(t *testing.T, x int, y int, opponents []int, expect int) {
	res := solve(x, y, opponents)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 2, 10
	opponents := []int{3, 1, 9, 2, 5, 20, 8}
	expect := 20
	runSample(t, x, y, opponents, expect)
}

func TestSample2(t *testing.T) {
	x, y := 1, 10
	opponents := []int{3, 1, 9, 2, 5, 20, 8}
	expect := -1
	runSample(t, x, y, opponents, expect)
}
func TestSample3(t *testing.T) {
	x, y := 10, 12
	opponents := []int{100, 1, 200, 11, 300}
	expect := 2
	runSample(t, x, y, opponents, expect)
}
