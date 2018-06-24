package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q int, A []byte, B []int, C []interface{}, expect []string) {
	res := solve(Q, A, B, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := 5
	A := []byte("++?+?")
	B := []int{0, 1, 1, 2, 1}
	C := []interface{}{"ab", "c", 3, "dd", 5}
	expect := []string{"acb", "acddb"}
	runSample(t, Q, A, B, C, expect)
}
