package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, G []int, expect []int) {
	res := solve(n, A, G)

	if res != (len(expect) > 0) {
		t.Fatalf("Sample expect %v, but got %t", expect, res)
	}
	if res && !reflect.DeepEqual(A, expect) {
		t.Fatalf("Sample result %v, not correct, %v", A, expect)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{-1, 0, -1}
	G := []int{0, 3, 0, 0}
	expect := []int{1, 0, 2}
	runSample(t, n, A, G, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{2, 2, 1}
	G := []int{3, 0, 0, 0}
	//expect :=
	runSample(t, n, A, G, nil)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{-1, 2, -1}
	G := []int{1, 1, 1, 0}
	expect := []int{0, 2, 0}
	runSample(t, n, A, G, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	A := []int{-1, 2, 2}
	G := []int{0, 1, 1, 0}
	//expect :=
	runSample(t, n, A, G, nil)
}
