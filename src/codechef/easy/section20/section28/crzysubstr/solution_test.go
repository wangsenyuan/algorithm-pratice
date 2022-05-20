package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S string, expect []int) {
	res := solve(n, S)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := "aaaa"
	expect := []int{4, 4, 4, 4, 4}
	runSample(t, n, S, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	S := "babbaab"
	expect := []int{2, 3, 4, 4, 4, 4, 4, 4}
	runSample(t, n, S, expect)
}
