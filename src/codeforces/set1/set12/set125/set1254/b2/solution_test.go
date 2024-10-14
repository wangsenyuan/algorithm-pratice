package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 8, 5}
	expect := 9
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 10, 2, 1, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 5, 15, 10}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{61, 28, 67, 53, 13, 6, 70, 5, 79, 82, 60, 60, 84, 17, 80, 25, 82, 82, 69, 76, 81, 43, 58, 86, 18, 78, 4, 25, 8, 30, 33, 87, 91, 18, 90, 26, 62, 11, 28, 66, 9, 33, 58, 66, 47, 48, 80, 38, 25, 57, 4, 84, 79, 71, 54, 84, 63, 32, 97, 62, 26, 68, 5, 69, 54, 93, 25, 26, 100, 73, 24, 94, 80, 39, 30, 45, 95, 80, 0, 29, 57, 98, 92, 15, 17, 76, 69, 11, 57, 56, 48, 10, 28, 7, 63, 66, 53, 58, 12, 58}
	expect := 68
	runSample(t, a, expect)
}
