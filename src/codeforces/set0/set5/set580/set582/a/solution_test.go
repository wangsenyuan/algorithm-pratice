package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, g []int, expect []int) {
	res := solve(n, g)

	sort.Ints(expect)
	sort.Ints(res)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	g := []int{2, 1, 2, 3, 4, 3, 2, 6, 1, 1, 2, 2, 1, 2, 3, 2}
	expect := []int{4, 3, 6, 2}
	runSample(t, n, g, expect)
}
