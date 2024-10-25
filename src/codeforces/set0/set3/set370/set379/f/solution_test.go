package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, steps []int, expect []int32) {
	res := solve(steps)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	steps := []int{2, 3, 4, 8, 5}
	expect := []int32{3, 4, 4, 5, 6}
	runSample(t, steps, expect)
}
