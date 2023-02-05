package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect [][]int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0}
	Q := [][]int{
		{1, 1},
	}
	expect := [][]int{
		{1, 1},
	}
	runSample(t, A, Q, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 10}
	Q := [][]int{
		{1, 2},
	}
	expect := [][]int{
		{1, 1},
	}
	runSample(t, A, Q, expect)
}

func TestSample3(t *testing.T) {
	A := []int{0, 2, 4}
	Q := [][]int{
		{1, 3},
	}
	expect := [][]int{
		{2, 2},
	}
	runSample(t, A, Q, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 12, 8, 3}
	Q := [][]int{
		{1, 4},
	}
	expect := [][]int{
		{2, 3},
	}
	runSample(t, A, Q, expect)
}

func TestSample5(t *testing.T) {
	A := []int{21, 32, 32, 32, 10}
	Q := [][]int{
		{1, 5},
	}
	expect := [][]int{
		{2, 3},
	}
	runSample(t, A, Q, expect)
}

func TestSample6(t *testing.T) {
	A := []int{0, 1, 0, 1, 0, 1, 0}
	Q := [][]int{
		{1, 7},
	}
	expect := [][]int{
		{2, 4},
	}
	runSample(t, A, Q, expect)
}
