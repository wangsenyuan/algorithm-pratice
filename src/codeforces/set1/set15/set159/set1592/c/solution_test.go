package main

import "testing"

func runSample(t *testing.T, k int, A []int, E [][]int, expect bool) {
	res := solve(k, A, E)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	A := []int{1, 3}
	E := [][]int{
		{1, 2},
	}
	expect := false
	runSample(t, k, A, E, expect)
}

func TestSample2(t *testing.T) {
	k := 5
	A := []int{3, 3, 3, 3, 3}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
	}
	expect := true
	runSample(t, k, A, E, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	A := []int{1, 7, 2, 3, 5}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
		{5, 3},
	}
	expect := false
	runSample(t, k, A, E, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	A := []int{1, 6, 4, 1, 2}
	E := [][]int{
		{1, 2},
		{2, 3},
		{1, 4},
		{4, 5},
	}
	expect := true
	runSample(t, k, A, E, expect)
}

func TestSample5(t *testing.T) {
	k := 3
	A := []int{1, 7, 4}
	E := [][]int{
		{1, 2},
		{2, 3},
	}
	expect := false
	runSample(t, k, A, E, expect)
}
