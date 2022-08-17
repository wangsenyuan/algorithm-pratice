package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, X, Y int64, expect []int64) {
	res := solve(X, Y)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	var X, Y int64 = 1, 6
	expect := []int64{2, 3, 5}
	runSample(t, X, Y, expect)
}

func TestSample2(t *testing.T) {
	var X, Y int64 = 9, 15
	expect := []int64{2, 11, 13}
	runSample(t, X, Y, expect)
}
