package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, p []int, expect []int) {
	res := solve(p)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{3, 2, 4, 1, 5}
	expect := []int{1, 0, 1, 1, 0}
	runSample(t, p, expect)
}
