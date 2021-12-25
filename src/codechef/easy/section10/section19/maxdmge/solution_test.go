package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{29, 412, 671, 255, 912}
	expect := []int{28, 156, 159, 159, 144}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{166, 94, 184, 114, 1124, 298}
	expect := []int{6, 24, 48, 96, 96, 32}
	runSample(t, A, expect)
}
