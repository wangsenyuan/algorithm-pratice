package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, K int, N int, queries []int, expect []int) {
	res := solve(K, N, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", K, N, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	K, N := 5, 5
	queries := []int{1, 2, 3, 4, 5}
	expect := []int{1, 3, 2, 5, 4}
	runSample(t, K, N, queries, expect)
}

func TestSample2(t *testing.T) {
	K, N := 15, 4
	queries := []int{3, 4, 7, 10}
	expect := []int{2, 8, 13, 4}
	runSample(t, K, N, queries, expect)
}
