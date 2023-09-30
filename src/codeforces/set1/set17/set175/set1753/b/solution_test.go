package main

import "testing"

func runSample(t *testing.T, a []int, x int, expect bool) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 2, 2, 3, 3}
	x := 4
	expect := true
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 2, 2, 2, 2, 2, 1, 1}
	x := 3
	expect := true
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 7, 7, 7, 7, 7, 7}
	x := 8
	expect := false
	runSample(t, a, x, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 3, 2, 1, 4, 3, 2, 4, 3, 4}
	x := 5
	expect := false
	runSample(t, a, x, expect)
}

func TestSample5(t *testing.T) {
	a := []int{499999, 499999}
	x := 499999 + 1
	expect := false
	runSample(t, a, x, expect)
}
