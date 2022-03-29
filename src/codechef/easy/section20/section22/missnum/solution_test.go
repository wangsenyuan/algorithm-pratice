package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, X []int, expect bool) {
	res := solve(X)

	if res[0] >= 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else if expect {
		a, b := res[0], res[1]
		Y := make([]int, 4)
		Y[0] = a + b
		Y[1] = a - b
		Y[2] = a * b
		Y[3] = a / b
		sort.Ints(X)
		sort.Ints(Y)

		if !reflect.DeepEqual(X, Y) {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{-1, 72, 0, 17}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 4, 5, 6}
	expect := false
	runSample(t, A, expect)
}
