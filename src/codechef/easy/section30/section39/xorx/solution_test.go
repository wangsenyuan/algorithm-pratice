package main

import (
	"math/rand"
	"testing"
)

func runSample(t *testing.T, A []int) {
	update := func(i int, v int) int {
		i--
		A[i] ^= v
		res := A[0]
		for j := 1; j < len(A); j++ {
			if A[j] > res {
				res = A[j]
			}
		}
		return res
	}
	solve(len(A), update)

	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			t.Fatalf("Solution doesn't give the correct answer %v", A)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	n := 10
	r := rand.New(rand.NewSource(99))

	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = r.Intn(1 << 30)
	}

	runSample(t, A)
}
