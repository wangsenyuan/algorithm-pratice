package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, candies []int, queries []int, expect []int) {
	res := solve(candies, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	candies := []int{4, 3, 3, 1, 1, 4, 5, 9}
	queries := []int{1, 10, 50, 14, 15, 22, 30}
	expect := []int{1, 2, -1, 2, 3, 4, 8}
	runSample(t, candies, queries, expect)
}
