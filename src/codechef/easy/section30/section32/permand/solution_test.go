package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, expect bool) {
	ok, A, B := solve(n)

	if ok != expect {
		t.Errorf("Sample expect %t, but got %t", expect, ok)
	} else if ok {
		for i := 0; i < n; i++ {
			if A[i]&B[i] != 0 {
				t.Fatalf("Sample result %v %v, not correct", A, B)
			}
		}
		sort.Ints(B)

		if !reflect.DeepEqual(A, B) {
			t.Errorf("Sample result %v %v, not correct", A, B)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, true)
}
