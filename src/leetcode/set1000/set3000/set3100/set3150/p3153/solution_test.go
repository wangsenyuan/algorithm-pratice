package p3153

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, queries [][]int, expect []bool) {
	res := isArraySpecial(a, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{3, 4, 1, 2, 6}
	queries := [][]int{{0, 4}}
	expect := []bool{false}
	runSample(t, nums, queries, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{4, 3, 1, 6}
	queries := [][]int{{0, 2}, {2, 3}}
	expect := []bool{false, true}
	runSample(t, nums, queries, expect)
}
