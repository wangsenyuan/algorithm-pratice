package main

import "testing"

func runSample(t *testing.T, n int, A []int) {
	B := make([]int, n)
	copy(B, A)
	res := solve(n, B)

	if len(res) > 2*n {
		t.Errorf("Sample result %v, exceed 2 * n", res)
		return
	}

	for i := 0; i < len(res); i++ {
		x := res[i] - 1
		A[x] = findMex(A)
	}

	if !sorted(A) {
		t.Errorf("Sample after operations %v, not sorted %v", res, A)
	}
}

func findMex(arr []int) int {
	mem := make(map[int]bool)

	for i := 0; i < len(arr); i++ {
		mem[arr[i]] = true
	}
	var mex int

	for mem[mex] {
		mex++
	}

	return mex
}

func TestSample1(t *testing.T) {
	arr := []int{2, 2, 3}
	runSample(t, len(arr), arr)
}

func TestSample2(t *testing.T) {
	arr := []int{2, 1, 0}
	runSample(t, len(arr), arr)
}

func TestSample3(t *testing.T) {
	arr := []int{0, 7, 3, 1, 3, 7, 7}
	runSample(t, len(arr), arr)
}

func TestSample4(t *testing.T) {
	arr := []int{2, 0, 1, 1, 2, 4, 4, 2, 0}
	runSample(t, len(arr), arr)
}
