package p985

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, queries [][]int, expect []int) {
	res := sumEvenAfterQueries(A, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", A, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	expect := []int{8, 6, 2, 4}
	runSample(t, A, queries, expect)
}
