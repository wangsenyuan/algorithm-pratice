package main

import "testing"

func runSample(t *testing.T, A []int, expect int64) {
	all, res := solve(len(A), A)

	if all != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, all)
	}

	var tmp int64

	for i := 1; i < len(res); i++ {
		tmp += 2 * int64(abs(res[i]-res[0])) * int64(A[i-1])
	}

	if tmp != expect {
		t.Fatalf("Sample result %v, got wrong answer => %d", res, tmp)
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func TestSample1(t *testing.T) {
	A := []int{3, 8, 10, 6, 1}
	var expect int64 = 78
	runSample(t, A, expect)
}
