package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, x int, expect []int) {
	res := solve(m, x)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, x := 2, 3
	expect := []int{1, 1, 1}
	runSample(t, m, x, expect)
}
