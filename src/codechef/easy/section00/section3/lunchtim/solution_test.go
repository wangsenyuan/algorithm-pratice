package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, H []int, expect []int) {
	res := solve(len(H), H)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	H := []int{1, 2, 2, 3, 2}
	expect := []int{0, 1, 1, 0, 0}
	runSample(t, H, expect)
}
