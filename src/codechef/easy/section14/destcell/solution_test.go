package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, expect []int) {
	res := solve(n, m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", n, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 3
	expect := []int{6, 4, 3, 3, 2, 1}
	runSample(t, n, m, expect)
}
