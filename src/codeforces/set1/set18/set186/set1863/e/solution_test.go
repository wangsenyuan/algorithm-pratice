package main

import "testing"

func runSample(t *testing.T, k int, h []int, deps [][]int, expect int) {
	res := solve(k, h, deps)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	h := []int{12, 16, 18, 12}
	k := 24
	deps := [][]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 4},
	}
	expect := 24
	runSample(t, k, h, deps, expect)
}

func TestSample2(t *testing.T) {
	h := []int{2, 6, 5, 9}
	k := 10
	deps := [][]int{
		{1, 4},
		{2, 4},
		{3, 4},
	}
	expect := 7
	runSample(t, k, h, deps, expect)
}

func TestSample3(t *testing.T) {
	h := []int{5, 5}
	k := 10
	deps := [][]int{
		{1, 2},
	}
	expect := 0
	runSample(t, k, h, deps, expect)
}

func TestSample4(t *testing.T) {
	h := []int{8, 800, 555, 35, 35}
	k := 1000
	//deps := [][]int{
	//	{1, 2},
	//}
	expect := 480
	runSample(t, k, h, nil, expect)
}

func TestSample5(t *testing.T) {
	h := []int{3, 2, 5, 4, 7}
	k := 10
	//deps := [][]int{
	//	{1, 2},
	//}
	expect := 5
	runSample(t, k, h, nil, expect)
}

func TestSample6(t *testing.T) {
	h := []int{4, 3, 2}
	k := 5
	deps := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := 8
	runSample(t, k, h, deps, expect)
}

func TestSample7(t *testing.T) {
	h := []int{3, 22, 15, 0, 20}
	k := 24
	deps := [][]int{
		{4, 5},
		{3, 4},
		{3, 5},
		{1, 2},
		{1, 5},
	}
	expect := 31
	runSample(t, k, h, deps, expect)
}
