package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := solve(len(s), s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "101"
	expect := []int{4, 5, 6, 7}
	runSample(t, s, expect)
}
