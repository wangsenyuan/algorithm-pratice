package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, pos []int, x string, expect []int) {
	res := solve(s, pos, []byte(x))
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcabcabc"
	pos := []int{1, 1, 2, 3, 4, 5, 8, 9, 1, 4}
	x := "abcabcabca"
	expect := []int{3, 2, 2, 2, 1, 2, 1, 1, 1, 0}
	runSample(t, s, pos, x, expect)
}
