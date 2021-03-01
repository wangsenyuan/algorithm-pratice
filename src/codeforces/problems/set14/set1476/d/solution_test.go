package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := solve(len(s), s)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "LRRRLL"
	expect := []int{1, 3, 2, 3, 1, 3, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "LRL"
	expect := []int{1, 4, 1, 4}
	runSample(t, s, expect)
}
