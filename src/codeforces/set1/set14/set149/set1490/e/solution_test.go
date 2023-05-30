package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 3, 2}
	expect := []int{1, 2}
	runSample(t, A, expect)
}
