package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m, n int, expect []int) {
	res := solve(m, n)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", m, n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 10, []int{2, 3, 5, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 5, []int{3, 5})
}
