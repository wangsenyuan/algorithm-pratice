package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, Q int, A []byte, B []int, C []int, expect []int) {
	res := solve(Q, A, B, C)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %v, expect %v, but got %v", Q, A, B, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	Q := 10
	A := []byte{'+', '-', '+', '?', '?', '?', '+', '?', '?', '-'}
	B := []int{8, 1, 3, 8, 8, 12, 6, 7, 4, 9}
	C := []int{1, 0, 9, 4, 8, 0, 5, 8, 5, 0}
	expect := []int{10, 8, 2, 4, 11}
	runSample(t, Q, A, B, C, expect)
}
