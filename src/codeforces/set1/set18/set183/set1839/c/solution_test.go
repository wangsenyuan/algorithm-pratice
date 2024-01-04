package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	n := len(a)

	b := make([]int, n)

	for _, p := range res {
		for i := 0; i < p; i++ {
			b[i] ^= 1
		}
		for j := n - 1; j > p; j-- {
			b[j] = b[j-1]
		}
		b[p] = 0
	}
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Sample result %v, got wrong final result %v", res, b)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 0, 0, 0}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 0, 1, 1, 0}
	expect := true
	runSample(t, a, expect)
}
