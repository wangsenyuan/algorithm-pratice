package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	ask := func(i, j int) int {
		return arr[i-1] | arr[j-1]
	}
	res := solve(len(arr), ask)

	if !reflect.DeepEqual(res, arr) {
		t.Errorf("Sample %v, got wrong result %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{2, 3, 1, 0}
	runSample(t, arr)
}

func TestSample2(t *testing.T) {
	arr := []int{3, 2, 0, 1, 4}
	runSample(t, arr)
}
