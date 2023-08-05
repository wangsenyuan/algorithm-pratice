package main

import "testing"

func runSample(t *testing.T, A [][]int, expect []int) {
	res := solve(A)
	if (len(res) == 0) != (len(expect) == 0) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if len(expect) == 0 {
		return
	}
	X := createMat(expect)
	Y := createMat(res)

	a := process(A, X)
	b := process(A, Y)

	if a != b {
		t.Fatalf("Sample result %v, not correct", res)
	}
	a = process(X, A)
	b = process(Y, A)

	if a != b {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func createMat(y []int) [][]int {
	n := len(y)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = y[i] ^ y[j]
		}
	}
	return mat
}

func process(a, b [][]int) int {
	n := len(a)
	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 0 || b[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func TestSample1(t *testing.T) {
	A := [][]int{
		{0, 0, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{1, 1, 0, 0},
	}
	expect := []int{0, 0, 1, 1}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := [][]int{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{1, 1, 0, 0},
	}
	//expect := []int{0, 0, 1, 1}
	runSample(t, A, nil)
}

func TestSample3(t *testing.T) {
	A := [][]int{
		{0, 0},
		{0, 0},
	}
	expect := []int{0, 1}
	runSample(t, A, expect)
}
