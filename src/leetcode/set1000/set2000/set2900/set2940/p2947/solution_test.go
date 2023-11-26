package p2947

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, limit int, expect []int) {
	res := lexicographicallySmallestArray(a, limit)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 52, 38, 59, 71, 27, 31, 83, 88, 10}
	limit := 14
	expect := []int{4, 27, 31, 38, 52, 59, 71, 83, 88, 10}
	runSample(t, a, limit, expect)
}
