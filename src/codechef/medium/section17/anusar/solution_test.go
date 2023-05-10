package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, F []int, expect []int64) {
	res := solve(s, F)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaeddf"
	F := []int{1, 2, 3, 4}
	expect := []int64{21, 4, 0, 0}
	runSample(t, s, F, expect)
}
