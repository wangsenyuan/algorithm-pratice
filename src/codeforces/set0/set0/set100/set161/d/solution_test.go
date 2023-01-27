package main

import "testing"

func runSample(t *testing.T, n int, k int, E [][]int, expect int) {
	res := int(solve(n, k, E))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 2
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{2, 5},
	}
	expect := 4
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 5, 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := 2
	runSample(t, n, k, E, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 3
	E := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{2, 5},
	}
	expect := 2
	runSample(t, n, k, E, expect)
}

func TestSample4(t *testing.T) {
	n, k := 10, 3
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 4},
		{6, 5},
		{7, 1},
		{8, 6},
		{9, 2},
		{10, 6},
	}
	expect := 8
	runSample(t, n, k, E, expect)
}

func TestSample5(t *testing.T) {
	n, k := 50, 4
	E := [][]int{
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 4},
		{9, 4},
		{10, 5},
		{11, 5},
		{12, 6},
		{13, 6},
		{14, 7},
		{15, 7},
		{16, 8},
		{17, 8},
		{18, 9},
		{19, 9},
		{20, 10},
		{21, 10},
		{22, 11},
		{23, 11},
		{24, 12},
		{25, 12},
		{26, 13},
		{27, 13},
		{28, 14},
		{29, 14},
		{30, 15},
		{31, 15},
		{32, 16},
		{33, 16},
		{34, 17},
		{35, 17},
		{36, 18},
		{37, 18},
		{38, 19},
		{39, 19},
		{40, 20},
		{41, 20},
		{42, 21},
		{43, 21},
		{44, 22},
		{45, 22},
		{46, 23},
		{47, 23},
		{48, 24},
		{49, 24},
		{50, 25},
	}
	expect := 124
	runSample(t, n, k, E, expect)
}
