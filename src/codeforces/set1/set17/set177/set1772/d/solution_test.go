package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)
	if (expect < 0) != (res < 0) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	if expect >= 0 {
		B := make([]int, len(A))
		copy(B, A)
		for i := 0; i < len(B); i++ {
			B[i] = abs(B[i] - res)
		}
		if !sort.IntsAreSorted(B) {
			t.Fatalf("Sample result %d, gives wrong array %v", res, B)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	A := []int{5, 3, 3, 3, 5}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{5, 3, 4, 5}
	expect := -1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expect := 0
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{10, 5, 4, 3, 2, 1}
	expect := 42
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{3, 3, 1}
	expect := 2
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{42, 43, 42}
	expect := -1
	runSample(t, A, expect)
}

func TestSample7(t *testing.T) {
	A := []int{29613295, 52036613, 75100585, 78027446, 81409090, 73215}
	expect := 40741153
	runSample(t, A, expect)
}
