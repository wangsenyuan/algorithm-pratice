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
	n, m := 2, 2
	a := []int{1, 1, 2, 2}
	b := []int{1, 1, 2, 2}
	expect := []int{0, 4}
	runSample(t, n, m, a, b, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 4
	a := []int{2, 2, 3, 2}
	b := []int{3, 1, 4, 3}
	expect := []int{3, 9}
	runSample(t, n, m, a, b, expect)
}
func TestSample3(t *testing.T) {
	n, m := 1, 5
	a := []int{1, 1, 5, 1}
	b := []int{3, 1, 5, 1}
	expect := []int{2, 3}
	runSample(t, n, m, a, b, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 4
	a := []int{1, 1, 4, 2}
	b := []int{1, 3, 4, 4}
	expect := []int{8, 8}
	runSample(t, n, m, a, b, expect)
}

func TestSample5(t *testing.T) {
	n, m := 3, 4
	a := []int{1, 2, 4, 2}
	b := []int{2, 1, 3, 3}
	expect := []int{4, 8}
	runSample(t, n, m, a, b, expect)
}
