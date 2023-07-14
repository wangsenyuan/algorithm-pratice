package main

import "testing"

func runSample(t *testing.T, r int, celebrities [][]int, expect int) {
	res := solve(r, celebrities)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := 10
	celebrities := [][]int{
		{11, 6, 8},
	}
	expect := 0
	runSample(t, r, celebrities, expect)
}

func TestSample2(t *testing.T) {
	r := 6
	celebrities := [][]int{
		{1, 2, 6},
		{7, 5, 1},
		{8, 5, 5},
		{10, 3, 1},
		{12, 4, 4},
		{13, 6, 2},
		{17, 6, 6},
		{20, 1, 4},
		{21, 5, 4},
	}
	expect := 4
	runSample(t, r, celebrities, expect)
}

func TestSample3(t *testing.T) {
	r := 10
	celebrities := [][]int{
		{1, 2, 1},
		{5, 10, 9},
		{13, 8, 8},
		{15, 9, 9},
	}
	expect := 1
	runSample(t, r, celebrities, expect)
}

func TestSample4(t *testing.T) {
	r := 500
	celebrities := [][]int{
		{69, 477, 122},
		{73, 186, 235},
		{341, 101, 145},
		{372, 77, 497},
		{390, 117, 440},
		{494, 471, 37},
		{522, 300, 498},
		{682, 149, 379},
		{821, 486, 359},
		{855, 157, 386},
	}
	expect := 3
	runSample(t, r, celebrities, expect)
}
