package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 2, 1, 2, 2, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 2, 3, 3, 4, 5, 6, 6, 6}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 3, 2, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{9, 7, 4, 6, 9, 9, 10}
	expect := true
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1000000000}
	expect := true
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{9, 4, 4, 5, 9, 4, 9, 10}
	expect := false
	runSample(t, a, expect)
}
