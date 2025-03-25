package main

import "testing"

func runSample(t *testing.T, s int, a []int, b []int, expect int) {
	res := solve(s, a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := 15
	a := []int{9, 2, 6, 5}
	b := []int{3, 5, 8, 9}
	expect := 6
	runSample(t, s, a, b, expect)
}

func TestSample2(t *testing.T) {
	s := 100
	a := []int{39, 9, 36, 94, 40, 26, 12, 26, 28, 66, 73, 85, 62, 5, 20}
	b := []int{0, 0, 7, 7, 0, 5, 5, 0, 7, 9, 9, 4, 2, 5, 2}
	expect := 119
	runSample(t, s, a, b, expect)
}
