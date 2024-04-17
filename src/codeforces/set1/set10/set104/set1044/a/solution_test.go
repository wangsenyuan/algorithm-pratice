package main

import "testing"

func runSample(t *testing.T, vs []int, hs [][]int, expect int) {
	res := solve(vs, hs)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	vs := []int{6, 8}
	hs := [][]int{
		{1, 5, 6},
		{1, 9, 4},
		{2, 4, 2},
	}
	expect := 1
	runSample(t, vs, hs, expect)
}

func TestSample2(t *testing.T) {
	vs := []int{4}
	hs := [][]int{
		{1, 5, 3},
		{1, 9, 4},
		{4, 6, 6},
	}
	expect := 1
	runSample(t, vs, hs, expect)
}

func TestSample3(t *testing.T) {
	//vs := []int{4}
	hs := [][]int{
		{1, 1000000000, 4},
		{1, 1000000000, 2},
	}
	expect := 2
	runSample(t, nil, hs, expect)
}

func TestSample4(t *testing.T) {
	vs := []int{4, 6}
	hs := [][]int{
		{1, 4, 3},
		{1, 5, 2},
		{1, 6, 5},
	}
	expect := 2
	runSample(t, vs, hs, expect)
}

func TestSample5(t *testing.T) {
	vs := []int{1, 2, 3, 4}
	hs := [][]int{
		{1, 1000000000, 1},
		{1, 1000000000, 2},
		{1, 1000000000, 3},
		{1, 1000000000, 4},
		{1, 1000000000, 5},
		{1, 1000000000, 6},
		{1, 1000000000, 7},
	}
	expect := 7
	runSample(t, vs, hs, expect)
}
