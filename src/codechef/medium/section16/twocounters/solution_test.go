package main

import "testing"

func runSample(t *testing.T, n int, T []int, C []int, expect int) {
	res := solve(n, T, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	T := []int{1, 2, 5}
	C := []int{1, 2, 1}
	expect := 2
	runSample(t, n, T, C, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	T := []int{1, 2, 3, 4, 5}
	C := []int{1, 1, 1, 1, 1}
	expect := 5
	runSample(t, n, T, C, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	T := []int{1, 2, 3, 4, 5}
	C := []int{1, 2, 1, 2, 1}
	expect := 3
	runSample(t, n, T, C, expect)
}

func TestSample4(t *testing.T) {
	n := 1
	T := []int{1}
	C := []int{2}
	expect := 1
	runSample(t, n, T, C, expect)
}
