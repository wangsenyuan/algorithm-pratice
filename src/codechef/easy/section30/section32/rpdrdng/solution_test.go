package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, B []int, expect bool) {
	res := solve(n, B)

	if expect != (len(res) > 0) {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	}

	if expect {
		C := generate(res)

		sort.Ints(B)
		sort.Ints(C)

		if !reflect.DeepEqual(B, C) {
			t.Errorf("Sample result %v, not correct", res)
		}
		if !sort.IntsAreSorted(res) {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func generate(A []int) []int {
	n := len(A)
	var B []int
	for i := 0; i < n; i++ {
		j := i / 2
		B = append(B, A[j])
	}
	for i := n - 1; i >= 0; i-- {
		suf := A[i:]
		m := (len(suf) - 1) / 2
		B = append(B, suf[m])
	}
	return B
}

func TestSample1(t *testing.T) {
	n := 3
	B := []int{6, 5, 5, 7, 6, 6}
	expect := true
	runSample(t, n, B, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	B := []int{1, 2, 1, 4, 2, 4}
	expect := false
	runSample(t, n, B, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	A := make([]int, 10)
	for i := 1; i <= 10; i++ {
		A[i-1] = i
	}
	B := generate(A)

	expect := true

	runSample(t, n, B, expect)
}

func TestSample4(t *testing.T) {
	n := 11
	A := make([]int, n)
	for i := 1; i <= n; i++ {
		A[i-1] = i
	}
	B := generate(A)

	expect := true

	runSample(t, n, B, expect)
}
