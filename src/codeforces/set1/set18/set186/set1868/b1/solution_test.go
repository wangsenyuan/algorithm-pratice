package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 4, 7, 1, 5, 4}
	expect := false
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{20092043, 20092043}
	expect := true
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{9, 9, 8, 2, 4, 4, 3, 5, 1, 1, 1, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 12, 7, 16, 11, 12}
	expect := false
	runSample(t, a, expect)
}
