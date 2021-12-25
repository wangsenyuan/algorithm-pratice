package main

import "testing"

func runSample(t *testing.T, N, K, D int, X []int, L []int, expect int) {
	res := solve(N, K, D, X, L)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, D := 2, 10, 20
	X := []int{4, 7}
	L := []int{1, 2}
	expect := 10
	runSample(t, N, K, D, X, L, expect)
}

func TestSample2(t *testing.T) {
	N, K, D := 4, 15, 20
	X := []int{4, 6, 9, 13}
	L := []int{1, 2, 2, 1}
	expect := 13
	runSample(t, N, K, D, X, L, expect)
}

func TestSample3(t *testing.T) {
	N, K, D := 5, 10, 1
	X := []int{1, 3, 5, 7, 9}
	L := []int{1, 2, 1, 2, 1}
	expect := 10
	runSample(t, N, K, D, X, L, expect)
}

func TestSample4(t *testing.T) {
	N, K, D := 2, 10, 2
	X := []int{4, 5}
	L := []int{1, 2}
	expect := 5
	runSample(t, N, K, D, X, L, expect)
}
