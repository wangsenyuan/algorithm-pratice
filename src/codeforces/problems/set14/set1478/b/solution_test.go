package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, d int, A []int, expect []bool) {
	res := solve(d, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{24, 25, 27}
	d := 7
	expect := []bool{true, false, true}
	runSample(t, d, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	d := 7
	expect := []bool{true, true, false, true, true, true, true, true, true, false}
	runSample(t, d, A, expect)
}
