package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, n int, A []int) {
	ask := func(k int) int {
		var res int
		for i := 0; i < n; i++ {
			res += A[i] ^ k
		}
		return res
	}

	var expect int

	for i := 0; i < n; i++ {
		expect ^= A[i]
	}

	ans := solve(n, ask)

	if ans != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, ans)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 2, 3, 4}
	runSample(t, n, A)
}

func TestSample2(t *testing.T) {
	n := 100
	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = rand.Intn(1000000) + 1
	}
	rand.Shuffle(n, func(i, j int) {
		A[i], A[j] = A[j], A[i]
	})

	runSample(t, n, A)
}
