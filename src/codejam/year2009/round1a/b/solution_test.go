package main

import "testing"

func runSample(t *testing.T, N int, M int, G [][]int, expect int) {
	res := solve(N, M, G)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, M, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 1, 1
	G := [][]int{
		{3, 2, 10},
	}
	expect := 4
	runSample(t, N, M, G, expect)
}

func TestSample2(t *testing.T) {
	N, M := 1, 2
	G := [][]int{
		{1, 5, 3, 1, 5, 2},
	}
	expect := 7
	runSample(t, N, M, G, expect)
}

func TestSample3(t *testing.T) {
	N, M := 2, 2
	G := [][]int{
		{1, 1, 0, 10, 1, 6},
		{10, 1, 0, 1, 10, 10},
	}
	expect := 12
	runSample(t, N, M, G, expect)
}

func TestSample4(t *testing.T) {
	N, M := 2, 2
	G := [][]int{
		{2, 10, 2, 10, 6, 10},
		{8, 7, 10, 4, 3, 2},
	}
	expect := 8
	runSample(t, N, M, G, expect)
}

func TestSample5(t *testing.T) {
	N, M := 1, 1
	G := [][]int{
		{3, 3, 5},
	}
	expect := 3
	runSample(t, N, M, G, expect)

}
