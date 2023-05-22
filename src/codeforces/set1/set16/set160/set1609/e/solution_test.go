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
	s := "aaabccccc"
	pos := []int{4, 4, 2, 5, 1, 6, 5, 2, 1, 5, 6, 7}
	x := "abbabbcaaabb"
	expect := []int{0, 1, 2, 2, 1, 2, 1, 2, 2, 2, 2, 2}
	runSample(t, s, pos, x, expect)
}
