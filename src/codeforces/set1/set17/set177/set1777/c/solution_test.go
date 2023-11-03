package main

import "testing"

func runSample(t *testing.T, a []int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 7}
	m := 4
	expect := -1
	runSample(t, a, m, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 7, 2, 9}
	m := 2
	expect := 0
	runSample(t, a, m, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 4, 3, 5, 7}
	m := 7
	expect := 3
	runSample(t, a, m, expect)
}
