package main

import "testing"

func runSample(t *testing.T, X, Y []int, expect int) {
	res := solve(X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{10}
	Y := []int{1}
	expect := 0
	runSample(t, X, Y, expect)
}

func TestSample2(t *testing.T) {
	X := []int{1, 3, 6}
	Y := []int{1, 1, 1}
	expect := 2
	runSample(t, X, Y, expect)
}

func TestSample3(t *testing.T) {
	X := []int{1, 3, 6}
	Y := []int{1, 2, 1}
	expect := 4
	runSample(t, X, Y, expect)
}

func TestSample5(t *testing.T) {
	X := []int{1, 2, 3, 6, 12, 18, 216, 24, 432, 864}
	Y := []int{30, 40, 30, 50, 100, 500, 400, 999, 9999, 123}
	expect := 128248098
	runSample(t, X, Y, expect)
}

func TestSample6(t *testing.T) {
	X := []int{10}
	Y := []int{1000}
	expect := 23226275
	runSample(t, X, Y, expect)
}
