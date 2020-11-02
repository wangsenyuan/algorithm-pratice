package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect int) {
	X := make([]int, n)
	copy(X, A)

	res := solve(n, X)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	sort.Ints(A)
	sort.Ints(X)

	if !reflect.DeepEqual(A, X) {
		t.Errorf("Sample result %v not correct", X)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := []int{1, 3, 2, 2, 4, 5, 4}
	expect := 3
	runSample(t, n, A, expect)
}
