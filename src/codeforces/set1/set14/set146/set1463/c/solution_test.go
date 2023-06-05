package main

import "testing"

func runSample(t *testing.T, cmds [][]int, expect int) {
	res := solve(cmds)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cmds := [][]int{
		{1, 5},
		{3, 0},
		{6, 4},
	}
	expect := 1
	runSample(t, cmds, expect)
}

func TestSample2(t *testing.T) {
	cmds := [][]int{
		{1, 5},
		{2, 4},
		{10, -5},
	}
	expect := 2
	runSample(t, cmds, expect)
}

func TestSample3(t *testing.T) {
	cmds := [][]int{
		{2, -5},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
	}
	expect := 0
	runSample(t, cmds, expect)
}

func TestSample4(t *testing.T) {
	cmds := [][]int{
		{3, 3},
		{5, -3},
		{9, 2},
		{12, 0},
	}
	expect := 2
	runSample(t, cmds, expect)
}

func TestSample5(t *testing.T) {
	cmds := [][]int{
		{1, 1},
		{2, -6},
		{7, 2},
		{8, 3},
		{12, -9},
		{14, 2},
		{18, -1},
		{23, 9},
	}
	expect := 1
	runSample(t, cmds, expect)
}

func TestSample6(t *testing.T) {
	cmds := [][]int{
		{1, -4},
		{4, -7},
		{6, -1},
		{7, -3},
		{8, -7},
	}
	expect := 1
	runSample(t, cmds, expect)
}

func TestSample7(t *testing.T) {
	cmds := [][]int{
		{1, 2},
		{2, -2},
	}
	expect := 0
	runSample(t, cmds, expect)
}

func TestSample8(t *testing.T) {
	cmds := [][]int{
		{3, 10},
		{5, 5},
		{8, 0},
		{12, -4},
		{14, -7},
		{19, -5},
	}
	expect := 2
	runSample(t, cmds, expect)
}
