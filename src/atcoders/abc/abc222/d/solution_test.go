package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1}
	b := []int{2, 3}
	expect := 5
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 2}
	b := []int{2, 2, 2}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	expect := 978222082
	runSample(t, a, b, expect)
}
