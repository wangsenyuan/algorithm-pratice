package main

import "testing"

func runSample(t *testing.T, N, K, J int, events [][]int, expect int) {
	res := solve(N, K, J, events)

	if res != expect {
		t.Errorf("Sample %d %d %d %v, expect %d, but got %d", N, K, J, events, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K, J := 4, 1, 2
	events := [][]int{
		{1, 1},
		{1, 2},
		{1, 1},
		{1, 2},
	}
	expect := 3
	runSample(t, N, K, J, events, expect)
}

func TestSample2(t *testing.T) {
	N, K, J := 4, 2, 2
	events := [][]int{
		{1, 1},
		{1, 2},
		{1, 1},
		{1, 2},
	}
	expect := 2
	runSample(t, N, K, J, events, expect)
}

func TestSample3(t *testing.T) {
	N, K, J := 4, 2, 1
	events := [][]int{
		{1, 1},
		{1, 2},
		{1, 2},
		{1, 2},
	}
	expect := 4
	runSample(t, N, K, J, events, expect)
}

func TestSample4(t *testing.T) {
	N, K, J := 4, 1, 4
	events := [][]int{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 2},
	}
	expect := 2
	runSample(t, N, K, J, events, expect)
}
