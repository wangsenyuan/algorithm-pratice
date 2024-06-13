package main

import "testing"

func runSample(t *testing.T, n int, m int, A [][]int, B [][]int, expect bool) {
	res := solve(n, m, A, B)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	A := [][]int{
		{1},
	}
	B := [][]int{
		{1},
	}
	expect := true
	runSample(t, n, m, A, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{4, 3},
		{2, 1},
	}
	expect := true
	runSample(t, n, m, A, B, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 2
	A := [][]int{
		{1, 2},
		{3, 4},
	}
	B := [][]int{
		{4, 3},
		{1, 2},
	}
	expect := false
	runSample(t, n, m, A, B, expect)
}

func TestSample4(t *testing.T) {
	n, m := 3, 4
	A := [][]int{
		{1, 5, 9, 6},
		{12, 10, 4, 8},
		{7, 11, 3, 2},
	}
	B := [][]int{
		{1, 5, 9, 6},
		{12, 10, 4, 8},
		{7, 11, 3, 2},
	}
	expect := true
	runSample(t, n, m, A, B, expect)
}

func TestSample5(t *testing.T) {
	n, m := 3, 3
	A := [][]int{
		{1, 5, 9},
		{6, 4, 2},
		{3, 8, 7},
	}
	B := [][]int{
		{9, 5, 1},
		{2, 4, 6},
		{7, 8, 3},
	}
	expect := true
	runSample(t, n, m, A, B, expect)
}
