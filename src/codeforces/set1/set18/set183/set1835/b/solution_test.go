package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, k int, a []int, expect []int) {
	res := solve(m, k, a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, k := 6, 2
	a := []int{1, 4, 5}
	expect := []int{4, 2}
	runSample(t, m, k, a, expect)
}

func TestSample2(t *testing.T) {
	m, k := 7, 1
	a := []int{2, 4, 7, 3, 0, 1, 6}
	expect := []int{1, 5}
	runSample(t, m, k, a, expect)
}
