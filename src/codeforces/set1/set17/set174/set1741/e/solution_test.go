package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 2, 3, 1, 3, 2, 2, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12, 1, 2, 7, 5}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 7, 8, 9, 10, 3}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 8, 6, 2}
	expect := false
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{3, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1}
	expect := false
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{4, 6, 2, 1, 9, 4, 9, 3, 4, 2}
	expect := true
	runSample(t, a, expect)
}
