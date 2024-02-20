package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect []int) {
	res := solve(a, b)
	x := getLis(expect)
	y := getLis(res)

	if len(x) != len(y) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 4}
	b := []int{5}
	expect := []int{6, 5, 4}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 7, 2, 4, 5}
	b := []int{5, 4, 1, 2, 7}
	expect := []int{1, 1, 7, 7, 2, 2, 4, 4, 5, 5}
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{9, 8, 7, 7, 6, 5, 4, 3, 2, 1}
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4}
	expect := []int{1, 3, 5, 2, 4}
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := []int{1, 9, 2, 3, 8, 1, 4, 7, 2, 9}
	b := []int{7, 8, 5, 4, 6}
	expect := []int{1, 9, 2, 3, 8, 8, 1, 4, 4, 7, 7, 2, 9, 6, 5}
	runSample(t, a, b, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1}
	b := []int{777}
	expect := []int{777, 1, 1, 1, 1, 1, 1}
	runSample(t, a, b, expect)
}
