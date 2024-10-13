package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 5, 7, 6, 7, 5, 4, 4, 6, 4}
	expect := []int{1, 1, 1, 2, -1, 1, 1, 3, 1, 1}
	runSample(t, a, expect)
}
