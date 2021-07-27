package main

import "testing"

func runSample(t *testing.T, N, D int, P, S []int, expect int) {
	res := int(solve(N, D, P, S))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, D := 2, 10
	P := []int{1, 2}
	S := []int{2, 1}
	expect := 3
	runSample(t, N, D, P, S, expect)
}

func TestSample2(t *testing.T) {
	N, D := 5, 7
	P := []int{1, 2, 3, 4, 5}
	S := []int{1, 2, 3, 4, 5}
	expect := 7
	runSample(t, N, D, P, S, expect)
}

func TestSample3(t *testing.T) {
	N, D := 5, 7
	P := []int{6, 6, 6, 6, 6}
	S := []int{5, 4, 3, 2, 1}
	expect := 5
	runSample(t, N, D, P, S, expect)
}

func TestSample4(t *testing.T) {
	N, D := 5, 5
	P := []int{1, 4, 5, 1, 3}
	S := []int{2, 1, 2, 2, 1}
	expect := 4
	runSample(t, N, D, P, S, expect)
}

func TestSample5(t *testing.T) {
	N, D := 5, 10
	P := []int{2, 3, 4, 5, 6}
	S := []int{2, 3, 3, 4, 1}
	expect := 7
	runSample(t, N, D, P, S, expect)
}
