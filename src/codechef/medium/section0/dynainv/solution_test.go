package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, m int, expect []int) {
	res := solve(n, A, m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %d, expect %v, but got %v", n, A, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 1
	A := []int{1, 2, 3, 4, 5}
	expect := []int{1}
	runSample(t, n, A, m, expect)
}
