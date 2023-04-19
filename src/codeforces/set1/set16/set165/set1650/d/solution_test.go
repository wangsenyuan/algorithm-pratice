package main

import (
	rand2 "math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int) {
	B := make([]int, len(A))
	copy(B, A)
	_, res := solve(A)

	C := make([]int, len(A))

	for i, x := range res {
		C[i] = i + 1
		if x > 0 {
			reverse(C[0 : i+1])
			reverse(C[0:x])
			reverse(C[x : i+1])
		}
	}

	if !reflect.DeepEqual(B, C) {
		t.Fatalf("Sample result %v, not giving the expect result %v, but got %v", res, B, C)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 5, 6, 1, 4}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := []int{3, 1, 2}
	runSample(t, A)
}

func TestSample3(t *testing.T) {
	A := []int{5, 8, 1, 3, 2, 6, 4, 7}
	runSample(t, A)
}

func TestSample4(t *testing.T) {
	n := 100
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i + 1
	}
	rand2.Shuffle(n, func(i, j int) {
		A[i], A[j] = A[j], A[i]
	})
	runSample(t, A)
}
