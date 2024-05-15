package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 8, 7, 2, 4}
	b := []int{9, 7, 2, 9, 3}
	expect := 646
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1000000}
	b := []int{1000000}
	expect := 757402647
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 3}
	b := []int{4, 2}
	expect := 20
	runSample(t, a, b, expect)
}
