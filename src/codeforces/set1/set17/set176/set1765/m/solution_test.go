package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 15
	expect := []int{5, 10}
	runSample(t, n, expect)
}
