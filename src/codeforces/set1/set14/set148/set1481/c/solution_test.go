package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, A, B, C []int, expect bool) {
	X := make([]int, n)
	copy(X, A)
	ok, res := solve(n, m, A, B, C)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if ok {
		for i := 0; i < m; i++ {
			X[res[i]-1] = C[i]
		}
		if !reflect.DeepEqual(X, B) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	A := []int{1}
	B := []int{1}
	C := []int{1}
	expect := true
	runSample(t, n, m, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 2
	A := []int{1, 2, 2, 1, 1}
	B := []int{1, 2, 2, 1, 1}
	C := []int{1, 2}
	expect := true
	runSample(t, n, m, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 2
	A := []int{1, 2, 2, 1, 1}
	B := []int{1, 2, 2, 1, 1}
	C := []int{1, 2}
	expect := true
	runSample(t, n, m, A, B, C, expect)
}

func TestSample4(t *testing.T) {
	n, m := 3, 3
	A := []int{2, 2, 2}
	B := []int{2, 2, 2}
	C := []int{2, 3, 2}
	expect := true
	runSample(t, n, m, A, B, C, expect)
}

func TestSample5(t *testing.T) {
	n, m := 10, 5
	A := []int{7, 3, 2, 1, 7, 9, 4, 2, 7, 9}
	B := []int{9, 9, 2, 1, 4, 9, 4, 2, 3, 9}
	C := []int{9, 9, 7, 4, 3}
	expect := true
	runSample(t, n, m, A, B, C, expect)
}

func TestSample6(t *testing.T) {
	n, m := 5, 2
	A := []int{1, 2, 2, 1, 1}
	B := []int{1, 2, 2, 1, 1}
	C := []int{3, 3}
	expect := false
	runSample(t, n, m, A, B, C, expect)
}

func TestSample7(t *testing.T) {
	n, m := 6, 4
	A := []int{3, 4, 2, 4, 1, 2}
	B := []int{2, 3, 1, 3, 1, 1}
	C := []int{2, 2, 3, 4}
	expect := false
	runSample(t, n, m, A, B, C, expect)
}
