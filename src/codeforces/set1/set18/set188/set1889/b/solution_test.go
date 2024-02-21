package main

import "testing"

func runSample(t *testing.T, c int, a []int, expect bool) {
	res := solve(c, a)

	if expect != res {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := 10
	a := []int{0, 20, 15, 10}
	expect := true
	runSample(t, c, a, expect)
}

func TestSample2(t *testing.T) {
	c := 1
	a := []int{1, 1}
	expect := true
	runSample(t, c, a, expect)
}

func TestSample3(t *testing.T) {
	c := 1
	a := []int{0, 1, 0, 4, 199}
	expect := true
	runSample(t, c, a, expect)
}

func TestSample4(t *testing.T) {
	c := 2
	a := []int{1, 1, 3, 1, 1}
	expect := false
	runSample(t, c, a, expect)
}

func TestSample5(t *testing.T) {
	c := 5
	a := []int{5, 6, 1, 10, 2}
	expect := false
	runSample(t, c, a, expect)
}

func TestSample6(t *testing.T) {
	c := 1000000
	a := []int{1000000000000, 1000000000000, 1000000000000, 1000000000000, 1000000000000}
	expect := true
	runSample(t, c, a, expect)
}

func TestSample7(t *testing.T) {
	c := 1
	a := []int{0, 0, 2}
	expect := false
	runSample(t, c, a, expect)
}
