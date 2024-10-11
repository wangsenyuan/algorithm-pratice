package main

import (
	"testing"
)

func runSample(t *testing.T, lands [][]int, expect int) {
	res := solve(lands)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lands := [][]int{
		{5, 5},
		{3, 4},
	}
	expect := 25
	runSample(t, lands, expect)
}

func TestSample2(t *testing.T) {
	lands := [][]int{
		{2, 5},
		{4, 3},
	}
	expect := 16
	runSample(t, lands, expect)
}

func TestSample3(t *testing.T) {
	lands := [][]int{
		{10, 1},
		{9, 8},
		{7, 6},
	}
	expect := 84
	runSample(t, lands, expect)
}
