package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{9, 6, 8, 4, 5, 2}
	b := []int{4, 1, 5, 6, 3, 1}
	expect := 32
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 3, 2}
	b := []int{3, 4, 9}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2}
	b := []int{1}
	expect := 1
	runSample(t, a, b, expect)
}
func TestSample4(t *testing.T) {
	a := []int{2, 3, 7, 10, 23, 28, 29, 50, 69, 135, 420, 1000}
	b := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	expect := 13824
	runSample(t, a, b, expect)
}
