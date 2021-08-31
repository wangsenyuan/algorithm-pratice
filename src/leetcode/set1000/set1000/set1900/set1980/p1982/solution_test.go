package p1982

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, n int, S []int) {
	res := recoverArray(n, S)
	X := make([]int, len(S))
	getAllSubSum(res, X)
	sort.Ints(S)
	sort.Ints(X)

	if !reflect.DeepEqual(S, X) {
		t.Errorf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	S := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	runSample(t, n, S)
}

func TestSample2(t *testing.T) {
	n := 2
	S := []int{0, 0, 0, 0}
	runSample(t, n, S)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []int{0, -1, 4, 5}
	S := make([]int, 1<<4)
	getAllSubSum(A, S)
	runSample(t, n, S)
}
