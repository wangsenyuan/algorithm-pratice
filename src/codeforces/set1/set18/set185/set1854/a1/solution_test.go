package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int) {
	arr := make([]int, len(a))
	copy(arr, a)
	res := solve(a)

	if len(res) > 50 {
		t.Fatalf("Sample result %v, takes too much steps", res)
	}

	for _, cur := range res {
		i, j := cur[0], cur[1]
		arr[i] += arr[j]
	}

	if !sort.IntsAreSorted(arr) {
		t.Fatalf("Sample result %v, not correct => %v", res, arr)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1}
	runSample(t, a)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, -10, 3}
	runSample(t, a)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 1, 1, 1}
	runSample(t, a)
}

func TestSample4(t *testing.T) {
	a := []int{0, 0, 0, 0, 0, 0, 0, 0}
	runSample(t, a)
}

func TestSample5(t *testing.T) {
	a := []int{1, 2, -4, 3, -10}
	runSample(t, a)
}

func TestSample6(t *testing.T) {
	a := []int{11, 12, 13, 14, 15, -15, -16, -17, -18, -19}
	runSample(t, a)
}

func TestSample7(t *testing.T) {
	a := []int{1, 9, 3, -4, -3, -2, -1}
	runSample(t, a)
}

func TestSample8(t *testing.T) {
	a := []int{10, 9, 8}
	runSample(t, a)
}

func TestSample9(t *testing.T) {
	a := []int{1, -14, 2, -10, 6, -5, 10, -13, 10, 7, -14, 19, -5, 19, 1, 18, -16, -7, 12, 8}
	runSample(t, a)
}

func TestSample10(t *testing.T) {
	a := []int{-15, -17, -13, 8, 14, -13, 10, -4, 11, -4, -16, -6, 15, -4, -2, 7, -9, 5, -5, 17}
	runSample(t, a)
}

func TestSample11(t *testing.T) {
	a := []int{-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20}
	runSample(t, a)
}
