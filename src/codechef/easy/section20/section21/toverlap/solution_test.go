package main

import (
	rand2 "math/rand"
	"testing"
)

func runSample(t *testing.T, A [][]int, B [][]int, expect int64) {
	res := solve(A, B)

	if res != expect {
		t.Errorf("Sample expect %d,  but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{1, 3},
		{2, 4},
	}
	var expect int64 = 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{1, 2},
	}
	B := [][]int{
		{3, 4},
	}
	var expect int64
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{5, 6},
		{6, 8},
	}
	B := [][]int{
		{5, 8},
	}
	var expect int64 = 3
	runSample(t, A, B, expect)
}

func createRandomSegs(n int) [][]int {

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		a, b := rand2.Intn(MAX_X), rand2.Intn(MAX_X)
		if a > b {
			a, b = b, a
		}
		res[i] = []int{a, b}
	}

	return res
}

func TestSolve1(t *testing.T) {
	A := createRandomSegs(10)
	B := createRandomSegs(20)
	expect := solve(A, B)
	res := solve1(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSolve2(t *testing.T) {
	A := createRandomSegs(10)
	B := createRandomSegs(20)
	expect := solve(A, B)
	res := solve2(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}
