package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, m int, s string, expect []int) {
	res := solve(a, m, s)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 4, 2}
	m := 6
	s := "LRRL"
	expect := []int{0, 2, 4, 1}
	runSample(t, a, m, s, expect)
}
