package p900

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, next []int, expect []int) {
	iter := Constructor(A)

	res := make([]int, len(next))
	for i, item := range next {
		res[i] = iter.Next(item)
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", A, next, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 8, 0, 9, 2, 5}
	next := []int{1, 1, 1, 1, 1, 1}
	expect := []int{8, 8, 8, 5, 5, -1}
	runSample(t, A, next, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 8, 0, 9, 2, 5}
	next := []int{2, 1, 1, 2}
	expect := []int{8, 8, 5, -1}
	runSample(t, A, next, expect)
}
