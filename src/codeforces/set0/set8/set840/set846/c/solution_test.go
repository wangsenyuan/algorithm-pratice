package main

import "testing"

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	n := len(A)
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + int64(A[i-1])
	}

	getResult := func(d0, d1, d2 int) int64 {
		tmp := sum[d0] - (sum[d1] - sum[d0]) + (sum[d2] - sum[d1]) - (sum[n] - sum[d2])
		return tmp
	}
	a := getResult(expect[0], expect[1], expect[2])
	b := getResult(res[0], res[1], res[2])
	// 2 * sum[d0] - 2 * sum[d1] + 2 * sum[d2] - sum[n]
	if a != b {
		t.Fatalf("Sample expect %d, but result %v, got %d", a, res, b)
	}
}

func TestSample1(t *testing.T) {
	A := []int{-1, 2, 3}
	expect := []int{0, 1, 3}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0, -1, 0}
	expect := []int{0, 0, 0}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{10000}
	expect := []int{1, 1, 1}
	runSample(t, A, expect)
}
