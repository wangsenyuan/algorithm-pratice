package main

import "testing"

func runSample(t *testing.T, n, m int, A []int, B []int) {
	res := solve(n, m, A, B)
	nums := make(map[int]bool)

	for _, row := range res {
		nums[A[row[0]]+B[row[1]]] = true
	}
	if len(nums) != n+m-1 {
		t.Error("Sample result is wrong")
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	A := []int{10, 1, 100}
	B := []int{4, 3}
	runSample(t, n, m, A, B)
}
