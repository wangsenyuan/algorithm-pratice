package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int64) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{10, 21, 32, 43, 54}
	expect := []int64{160, 150, 129, 97, 54}
	runSample(t, A, expect)
}
