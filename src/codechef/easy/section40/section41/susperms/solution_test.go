package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, m int, expect []int) {
	res := solve(k, m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, m := 8, 269696969
	expect := []int{2, 16, 576, 34560, 3110400, 142258231, 78756849, 13872609}
	runSample(t, k, m, expect)
}
