package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, K int, G []int, L [][]int, expect []int64) {
	res := solve(K, G, L)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	G := []int{-4, -2, 5, 2}
	L := [][]int{
		{1, 3, 0},
		{1, 4, - 3},
	}
	K := 3
	expect := []int64{7, 5, 5}
	runSample(t, K, G, L, expect)
}
