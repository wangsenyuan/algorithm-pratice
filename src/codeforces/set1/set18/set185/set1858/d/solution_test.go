package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, k int, expect []int) {
	res := solve(s, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "111"
	k := 0
	expect := []int{3, 3, 3}
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "0110"
	k := 1
	expect := []int{4, 5, 7, 9}
	runSample(t, s, k, expect)
}
