package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, x int, expect []int) {
	res := solve(n, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample(t *testing.T) {
	n, x := 12, 3
	expect := []int{3, 2, 6, 4, 5, 12, 7, 8, 9, 10, 11, 1}
	runSample(t, n, x, expect)
}
