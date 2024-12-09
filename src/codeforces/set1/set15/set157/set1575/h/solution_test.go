package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, x string, expect []int) {
	res := solve(s, x)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "100101011"
	x := "101"
	expect := []int{1, 1, 0, 1, 6, -1, -1, -1}
	runSample(t, s, x, expect)
}
