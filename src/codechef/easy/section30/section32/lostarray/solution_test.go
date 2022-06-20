package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, B []int) {
	A := solve(n, B)
	R := make([]int, n+1)
	for i := 0; i < n; i++ {
		R[n] ^= A[i]
	}
	for i := 0; i < n; i++ {
		R[i] = R[n] ^ A[i]
	}
	sort.Ints(B)
	sort.Ints(R)

	if !reflect.DeepEqual(R, B) {
		t.Errorf("Sample result %v, not correct", A)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	B := []int{2, 3, 0, 1}
	runSample(t, n, B)
}

func TestSample2(t *testing.T) {
	n := 5
	B := []int{27, 15, 23, 30, 29, 31}
	runSample(t, n, B)
}
