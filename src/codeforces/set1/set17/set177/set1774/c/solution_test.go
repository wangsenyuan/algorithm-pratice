package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, s string, expect []int) {
	res := solve(n, s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	s := "001"
	expect := []int{1, 1, 3}
	runSample(t, n, s, expect)
}
