package main

import "testing"

func runSample(t *testing.T, king []int, rocks [][]int, expect bool) {
	res := solve(king, rocks)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	king := []int{1, 1}
	rocks := [][]int{
		{8, 2},
		{7, 3},
	}
	expect := true

	runSample(t, king, rocks, expect)
}

func TestSample2(t *testing.T) {
	king := []int{1, 8}
	rocks := [][]int{
		{8, 2},
		{8, 3},
	}
	expect := false

	runSample(t, king, rocks, expect)
}

func TestSample3(t *testing.T) {
	king := []int{1, 1}
	rocks := [][]int{
		{8, 2},
		{2, 3},
	}
	expect := false

	runSample(t, king, rocks, expect)
}

func TestSample4(t *testing.T) {
	king := []int{1, 1}
	rocks := [][]int{
		{8, 8},
		{8, 2},
	}
	expect := false

	runSample(t, king, rocks, expect)
}
