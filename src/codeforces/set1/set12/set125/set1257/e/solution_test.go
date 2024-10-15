package main

import "testing"

func runSample(t *testing.T, a []int, b []int, c []int, expect int) {
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1}
	b := []int{4}
	c := []int{2, 5}
	expect := 1
	runSample(t, a, b, c, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 1}
	b := []int{5, 4}
	c := []int{6}
	expect := 0
	runSample(t, a, b, c, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 6}
	b := []int{4}
	c := []int{1, 2, 3}
	expect := 3
	runSample(t, a, b, c, expect)
}

func TestSample4(t *testing.T) {
	a := []int{6}
	b := []int{5, 1, 2, 4, 7}
	c := []int{3}
	expect := 2
	runSample(t, a, b, c, expect)
}
