package main

import (
	"math/rand"
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int) {
	arr := make([]int, len(A))
	copy(arr, A)
	res := solve(A)
	if len(res) > len(A) {
		t.Fatalf("Sample expect less than %d operations, but got %d", len(A), len(res))
	}
	for _, op := range res {
		i, j := op[0], op[1]
		i--
		j--
		if (arr[i]+arr[j])%2 == 0 {
			arr[i] = arr[j]
		} else {
			arr[j] = arr[i]
		}
	}

	if !sort.IntsAreSorted(arr) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1000000000, 3, 0, 5}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	n := 10000
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i + 1
	}
	rand.Shuffle(len(A), func(i, j int) {
		A[i], A[j] = A[j], A[i]
	})
	runSample(t, A)
}
