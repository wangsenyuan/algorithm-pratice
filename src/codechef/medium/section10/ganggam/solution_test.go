package main

import "testing"

func runSample(t *testing.T, n, X int, S, E []int, expect int) {
	res := solve(n, X, S, E)

	if res != expect {
		t.Errorf("sample %d, %v %v, expect %d, but got %d", n, S, E, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	X := 2
	S := []int{1, 3, 4}
	E := []int{3, 5, 6}
	expect := 1
	runSample(t, n, X, S, E, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	X := 3
	S := []int{1, 2, 3, 4}
	E := []int{8, 9, 10, 11}
	expect := 2
	runSample(t, n, X, S, E, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	X := 3
	S := []int{1, 4, 8, 12}
	E := []int{5, 9, 13, 17}
	expect := 0
	runSample(t, n, X, S, E, expect)
}
