package main

import "testing"

func runSample(t *testing.T, a []int, b []int, c []int, d []int, expect int) {
	res := solve(a, b, c, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 5, 7}
	b := []int{2, 1, 1, 2}
	c := []int{3, 7}
	d := []int{1, 1}
	expect := 8
	runSample(t, a, b, c, d, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1299721, 1999993}
	b := []int{100000, 265}
	c := []int{1299721, 1999993}
	d := []int{100000, 265}
	expect := 1
	runSample(t, a, b, c, d, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 5}
	b := []int{1, 1}
	c := []int{2, 3}
	d := []int{1, 1}
	expect := 0
	runSample(t, a, b, c, d, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1998041}
	b := []int{999}
	c := []int{1998041}
	d := []int{99997}
	expect := 0
	runSample(t, a, b, c, d, expect)
}
