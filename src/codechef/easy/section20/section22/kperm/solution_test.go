package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, r int, k int64, expect []int) {
	res := solve(n, r, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, r := 3, 1
	var k int64 = 2
	expect := []int{2, 1, 3}
	runSample(t, n, r, k, expect)
}
