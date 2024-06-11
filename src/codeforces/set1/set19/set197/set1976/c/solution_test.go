package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, a []int, b []int, expect []int) {
	res := solve(n, m, a, b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 0
	a := []int{2, 1}
	b := []int{1, 2}
	expect := []int{1, 2}
	runSample(t, n, m, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, m := 0, 2
	a := []int{4, 5, 5}
	b := []int{5, 4, 1}
	expect := []int{5, 6, 9}
	runSample(t, n, m, a, b, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 1
	a := []int{4, 3, 3, 4, 1}
	b := []int{5, 5, 4, 5, 2}
	expect := []int{13, 13, 13, 12, 15}
	runSample(t, n, m, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, m := 1, 2
	a := []int{2, 1, 5, 4}
	b := []int{5, 2, 3, 1}
	expect := []int{8, 11, 11, 12}
	runSample(t, n, m, a, b, expect)
}
