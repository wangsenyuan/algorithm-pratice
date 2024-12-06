package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, x string, expect []int) {
	res := solve(s, x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcdefa"
	x := "ddcb"
	expect := []int{2, 3, 3, 3}
	runSample(t, s, x, expect)
}
