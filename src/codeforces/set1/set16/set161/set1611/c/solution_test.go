package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect bool) {
	P := solve(A)

	if len(P) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, P)
	}
	if !expect {
		return
	}
	n := len(A)
	B := make([]int, 2*n)
	l, r := n, n

	L, R := 0, n-1
	for L < R {
		if P[L] < P[R] {
			l--
			B[l] = P[L]
			L++
		} else {
			B[r] = P[R]
			r++
			R--
		}
	}
	// L== R
	if A[0] == n {
		l--
		B[l] = n
	} else {
		B[r] = n
		r++
	}
	B = B[l:r]
	if !reflect.DeepEqual(B, A) {
		t.Fatalf("Sample result %v, giving wrong result %v", P, B)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 2, 4}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1}
	expect := true
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 5, 4, 2}
	expect := false
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 2, 1}
	expect := true
	runSample(t, A, expect)
}
