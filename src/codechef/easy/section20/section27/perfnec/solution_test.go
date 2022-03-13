package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, price []int, Q [][]int, expect []int64) {
	res := solve(price, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	price := []int{1, 2, 2, 3, 4}
	Q := [][]int{
		{2, 0},
		{1, -2},
		{3, 6},
	}
	expect := []int64{4, 1, 12}
	runSample(t, price, Q, expect)
}
