package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, m int, S string, Q []int, expect []int) {
	res := solve(n, m, []byte(S), Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	S := "010"
	Q := []int{2, 1, 3}
	expect := []int{4, 3, 2}
	runSample(t, n, m, S, Q, expect)
}
