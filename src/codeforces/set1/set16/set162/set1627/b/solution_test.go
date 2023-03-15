package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, expect []int) {
	res := solve(n, m)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 3
	expect := []int{3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5}
	runSample(t, n, m, expect)
}
