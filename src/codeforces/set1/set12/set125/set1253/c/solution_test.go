package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, m int, expect []int) {
	res := solve(a, m)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{6, 19, 3, 4, 4, 2, 6, 7, 8}
	m := 2
	expect := []int{2, 5, 11, 18, 30, 43, 62, 83, 121}
	runSample(t, a, m, expect)
}
