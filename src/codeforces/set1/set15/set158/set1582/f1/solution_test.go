package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	sort.Ints(expect)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 2, 4}
	expect := []int{0, 2, 4, 6}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 1, 7, 12, 5, 3, 2}
	expect := []int{0, 1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13}
	runSample(t, a, expect)
}
