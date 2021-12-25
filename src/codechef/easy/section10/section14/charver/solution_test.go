package main

import "testing"

func runSample(t *testing.T, C []int, A [][]int, expect bool) {
	res := solve(C, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{998244351, 998244348, 1}
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	expect := true
	runSample(t, C, A, expect)
}

func TestSample2(t *testing.T) {
	C := []int{998244351, 998244348, 1}
	A := [][]int{
		{1, 1},
		{1, 1},
	}
	expect := false
	runSample(t, C, A, expect)
}
