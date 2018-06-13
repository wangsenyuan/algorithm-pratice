package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, a, b int, expect []int) {
	res := solve(n, a, b)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d, %d, %d should give %v, but get %v", n, a, b, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 2, 1, []int{1, 2, 4, 5, 10, 11, 13, 14, 28, 29})
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 1, 1, []int{1, 2, 3, 4, 5})
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 4, 2, []int{1, 3, 4, 5})
}
