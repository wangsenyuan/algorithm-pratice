package main

import "testing"

func runSample(t *testing.T, H [][]int, M []string, k int, expect bool) {
	res := solve(H, M, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := [][]int{
		{7, 11, 3},
		{4, 2, 3},
		{0, 1, 15},
	}
	M := []string{
		"100",
		"010",
		"000",
	}
	k := 2
	expect := true
	runSample(t, H, M, k, expect)
}

func TestSample2(t *testing.T) {
	H := [][]int{
		{123, 413, 24, 233},
		{123, 42, 0, 216},
		{22, 1, 1, 53},
		{427, 763, 22, 6},
	}
	M := []string{
		"0101",
		"1111",
		"1010",
		"0101",
	}
	k := 3
	expect := false
	runSample(t, H, M, k, expect)
}

func TestSample3(t *testing.T) {
	H := [][]int{
		{0, 0},
		{2, 0},
	}
	M := []string{
		"01",
		"00",
	}
	k := 2
	expect := true
	runSample(t, H, M, k, expect)
}
