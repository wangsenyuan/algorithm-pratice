package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, parent []int, A []int64, expect []int64) {
	res := solve(n, parent, A)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("sample %d %v %v, expect %v but got %v", n, parent, A, expect, res)
	}
}

func TestSample(t *testing.T) {
	n := 8
	parent := []int{1, 1, 1, 1, 5, 8, 6}
	A := []int64{1, 2, 3, 4, 5, 15, 70, 10}
	expect := []int64{1, 3, 4, 5, 6, 21, 96, 26}
	runSample(t, n, parent, A, expect)
}
