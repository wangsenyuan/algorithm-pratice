package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	expect := []int{2, 1, 3}
	runSample(t, n, k, expect)
}
