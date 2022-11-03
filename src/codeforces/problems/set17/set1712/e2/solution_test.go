package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q [][]int, expect []int64) {
	res := solve(Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := [][]int{
		{1, 4},
		{3, 5},
		{8, 86},
		{68, 86},
		{6, 86868},
	}
	expect := []int64{3, 1, 78975, 969, 109229059713337}
	runSample(t, Q, expect)
}
