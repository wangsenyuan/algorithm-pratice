package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, K int64, arr []int, expect []int) {
	solve(n, K, arr)

	if !reflect.DeepEqual(arr, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, arr)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2}
	expect := []int{3, 1}
	runSample(t, 2, 2, arr, expect)
}
