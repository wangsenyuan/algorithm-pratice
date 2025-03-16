package p3488

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, nums []int, queries []int, expect []int) {
	res := solveQueries(nums, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 1, 4, 1, 3, 2}
	queries := []int{0, 3, 5}
	expect := []int{2, -1, 3}
	runSample(t, nums, queries, expect)
}
