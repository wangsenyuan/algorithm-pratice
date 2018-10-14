package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, q int, L []int, R []int, N []int, expect []int) {
	res, _ := solve(q, L, R, N)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %v %v, expect %v, but got %v", q, L, R, N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	q := 4
	L := []int{1, 6, 1, 6}
	R := []int{3, 9, 3, 9}
	N := []int{1, 3, 3, 4}
	expect := []int{1, 9, 3, -1}
	runSample(t, q, L, R, N, expect)
}
