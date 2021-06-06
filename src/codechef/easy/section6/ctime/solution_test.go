package main

import "testing"

func runSample(t *testing.T, N, K, F int, S [][]int, expect bool) {
	res := solve(N, K, F, S)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, F := 1, 10, 10
	S := [][]int{{0, 10}}
	expect := false
	runSample(t, N, K, F, S, expect)
}

func TestSample2(t *testing.T) {
	N, K, F := 2, 2, 10
	S := [][]int{{0, 5}, {7, 10}}
	expect := true
	runSample(t, N, K, F, S, expect)
}

func TestSample3(t *testing.T) {
	N, K, F := 2, 2, 100
	S := [][]int{{0, 5}, {5, 10}}
	expect := true
	runSample(t, N, K, F, S, expect)
}
