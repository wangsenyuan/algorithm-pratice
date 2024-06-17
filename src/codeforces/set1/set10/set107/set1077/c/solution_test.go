package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	sort.Ints(expect)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 5, 1, 2, 2}
	expect := []int{4, 1, 5}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 1, 2, 4, 3}
	runSample(t, a, nil)
}
